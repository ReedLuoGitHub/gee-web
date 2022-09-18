package gee

import (
	"fmt"
	"log"
)

type router struct {
	/* key -> method-pattern */
	handlers map[string]HandlerFunc
}

func newRoute() *router {
	return &router{handlers: make(map[string]HandlerFunc)}
}

func (r *router) addRoute(method, pattern string, handler HandlerFunc) {
	key := fmt.Sprintf("%s-%s", method, pattern)
	fmt.Println(key)
	r.handlers[key] = handler
}

// ServeHTTP 从router中拿到指定的handler进行处理
func (r *router) handler(ctx *Context) {
	if ctx.Path == "/favicon.ico" {
		return
	}
	key := fmt.Sprintf("%s-%s", ctx.Method, ctx.Path)
	handler, ok := r.handlers[key]
	if !ok {
		log.Println("404 not found: ", ctx.Method, ctx.Path)
		return
	}
	log.Println(ctx.Method, ctx.Path, ctx.StatusCode)
	handler(ctx)
}
