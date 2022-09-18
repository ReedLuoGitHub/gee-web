package gee

import (
	"fmt"
	"log"
	"net/http"
)

type HandlerFunc func(*Context)
type H map[string]interface{}

func init() {
	log.SetPrefix("[gee] ")
	log.SetFlags(log.LstdFlags | log.Lshortfile)
}

/* output */

func ErrPrint(info string) string {
	return fmt.Sprintf("%c[%d;%d;%dm%s(xxxx)%c[0m ", 0x1B, 41, 38, 2, info, 0x1B)
}

func NormalPrint(info string) string {
	return fmt.Sprintf(" %c[%d;%d;%dm%s(xxxx)%c[0m ", 0x1B, 42, 38, 5, info, 0x1B)
}

/* Engine */

type Engine struct {
	router *router
}

func New() *Engine {
	return &Engine{router: newRoute()}
}

// ServeHTTP 从router中拿到指定的handler进行处理
func (e Engine) ServeHTTP(w http.ResponseWriter, req *http.Request) {
	c := newContext(w, req)
	e.router.handler(c)
}

func (e Engine) Run(address string) {
	err := http.ListenAndServe(address, e)
	if err != nil {
		log.Fatal("start web server error: ", err)
	}
}

func (e Engine) addRoute(method, pattern string, handler HandlerFunc) {
	e.router.addRoute(method, pattern, handler)
}

func (e Engine) GET(pattern string, handler HandlerFunc) {
	e.addRoute("GET", pattern, handler)
}

func (e Engine) POST(pattern string, handler HandlerFunc) {
	e.addRoute("POST", pattern, handler)
}

func (e Engine) DELETE(pattern string, handler HandlerFunc) {
	e.addRoute("DELETE", pattern, handler)
}

func (e Engine) PUT(pattern string, handler HandlerFunc) {
	e.addRoute("PUT", pattern, handler)
}

func (e Engine) PATCH(pattern string, handler HandlerFunc) {
	e.addRoute("PATCH", pattern, handler)
}
