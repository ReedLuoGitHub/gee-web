package gee

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

type Context struct {
	Writer http.ResponseWriter
	Req    *http.Request

	// request
	Path   string
	Method string
	Params map[string]string

	// response
	StatusCode int

	// middleware
	handlers []HandlerFunc // 存放所有中间件以及路由对应的处理函数
	index    int           // 当前执行到哪个中间件，0开始
}

func newContext(w http.ResponseWriter, req *http.Request) *Context {
	return &Context{
		Writer: w,
		Req:    req,
		Path:   req.URL.Path,
		Method: req.Method,
		index:  -1,
	}
}

// Next 执行下一个中间件
func (ctx *Context) Next() {
	ctx.index++
	s := len(ctx.handlers)
	for ; ctx.index < s; ctx.index++ {
		ctx.handlers[ctx.index](ctx)
	}
}

func (ctx *Context) Param(key string) string {
	val, _ := ctx.Params[key]
	return val
}

func (ctx *Context) PostForm(key string) string {
	return ctx.Req.FormValue(key)
}

func (ctx *Context) Query(key string) string {
	return ctx.Req.URL.Query().Get(key)
}

func (ctx *Context) Status(code int) {
	ctx.StatusCode = code
	ctx.Writer.WriteHeader(code)
}

func (ctx *Context) SetHeader(key string, val string) {
	ctx.Writer.Header().Set(key, val)
}

func (ctx *Context) String(code int, format string, values ...interface{}) {
	ctx.SetHeader("Content-Type", "application/json")
	ctx.Status(code)
	_, err := ctx.Writer.Write([]byte(fmt.Sprintf(format, values...)))
	if err != nil {
		log.Println("response error: ", err)
	}
}

func (ctx *Context) JSON(code int, val interface{}) {
	ctx.SetHeader("Context-Type", "application/json")
	ctx.Status(code)
	err := json.NewEncoder(ctx.Writer).Encode(val)
	if err != nil {
		log.Println("response error: ", err)
	}
}

func (ctx *Context) Data(code int, val []byte) {
	ctx.Status(code)
	_, err := ctx.Writer.Write(val)
	if err != nil {
		log.Println("response error: ", err)
	}
}

func (ctx *Context) HTML(code int, html string) {
	ctx.SetHeader("Content-Type", "text/html")
	ctx.Status(code)
	_, err := ctx.Writer.Write([]byte(html))
	if err != nil {
		log.Println("response error: ", err)
	}
}
