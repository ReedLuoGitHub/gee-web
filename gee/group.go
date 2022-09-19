package gee

import (
	"log"
)

type RouterGroup struct {
	prefix      string        // 分组名
	middlewares []HandlerFunc // 该分组支持的中间件
	parent      *RouterGroup  // 用于group嵌套
	engine      *Engine       // 所有RouterGroup共享一个engine
}

func (g *RouterGroup) Group(prefix string) *RouterGroup {
	engine := g.engine
	newGroup := &RouterGroup{
		prefix: g.prefix + prefix,
		parent: g,
		engine: engine,
	}

	engine.groups = append(engine.groups, newGroup)
	return newGroup
}

func (g *RouterGroup) addRoute(method, comp string, handler HandlerFunc) {
	pattern := g.prefix + comp
	log.Printf("Route %4s- %s", method, pattern)
	g.engine.router.addRoute(method, pattern, handler)
}

func (g *RouterGroup) GET(pattern string, handler HandlerFunc) {
	g.addRoute("GET", pattern, handler)
}

func (g *RouterGroup) POST(pattern string, handler HandlerFunc) {
	g.addRoute("POST", pattern, handler)
}

func (g *RouterGroup) DELETE(pattern string, handler HandlerFunc) {
	g.addRoute("DELETE", pattern, handler)
}

func (g *RouterGroup) PUT(pattern string, handler HandlerFunc) {
	g.addRoute("PUT", pattern, handler)
}

func (g *RouterGroup) PATCH(pattern string, handler HandlerFunc) {
	g.addRoute("PATCH", pattern, handler)
}
