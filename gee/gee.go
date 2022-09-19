package gee

import (
	"log"
	"net/http"
	"strings"
)

type HandlerFunc func(*Context)
type H map[string]interface{}

func init() {
	log.SetPrefix("[gee] ")
	log.SetFlags(log.LstdFlags | log.Lshortfile)
}

/* Engine */

type Engine struct {
	router *router

	*RouterGroup                // engine拥有Group的所有功能
	groups       []*RouterGroup // 所有的分组信息
}

func New() *Engine {
	engine := &Engine{
		router: newRouter(),
	}
	engine.RouterGroup = &RouterGroup{engine: engine}
	engine.groups = []*RouterGroup{engine.RouterGroup}

	return engine
}

func Default() *Engine {
	engine := New()
	engine.Use(Logger())
	return engine
}

// ServeHTTP 从router中拿到指定的handler进行处理
func (e Engine) ServeHTTP(w http.ResponseWriter, req *http.Request) {
	var middlewares []HandlerFunc
	for _, group := range e.groups {
		// 通过URL前缀判断该请求适用于哪些中间件
		if strings.HasPrefix(req.URL.Path, group.prefix) {
			middlewares = append(middlewares, group.middlewares...)
		}
	}

	ctx := newContext(w, req)
	ctx.handlers = middlewares
	e.router.handle(ctx)
}

func (e Engine) Run(address string) {
	err := http.ListenAndServe(address, e)
	if err != nil {
		log.Fatal("start web server error: ", err)
	}
}
