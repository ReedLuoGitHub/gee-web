package gee

import (
	"fmt"
	"net/http"
	"strings"
)

type router struct {
	roots map[string]*node

	/* key -> method-pattern */
	handlers map[string]HandlerFunc
}

func newRouter() *router {
	return &router{
		roots:    make(map[string]*node),
		handlers: make(map[string]HandlerFunc),
	}
}

func (r *router) addRoute(method, pattern string, handler HandlerFunc) {
	// /hello/:name -> [hello :home]
	parts := parsePattern(pattern)

	// key -> GET-/home/:name
	key := fmt.Sprintf("%s-%s", method, pattern)

	// 如果没有该路由，则新建一个节点
	if _, ok := r.roots[method]; !ok {
		r.roots[method] = &node{}
	}

	// 将该节点插入到前缀树中
	r.roots[method].insert(pattern, parts, 0)

	// 新增一条 GET-/home/:name 和处理函数的映射
	r.handlers[key] = handler
}

// getRoute 解析通配符，返回通配符匹配到的字符
// /p/:lang/doc  /p/go/doc   ->   {"lang", "go"}
// /static/*filepath   /static/css/geektutu.css   ->   {"filepath", "css/geektutu.css"}
func (r *router) getRoute(method string, path string) (*node, map[string]string) {
	searchParts := parsePattern(path)
	params := make(map[string]string)

	// 拿到指定的Trie树节点
	root, ok := r.roots[method]
	if !ok {
		return nil, nil
	}

	n := root.search(searchParts, 0)
	if n != nil {
		parts := parsePattern(n.pattern)
		for idx, part := range parts {
			if part[0] == ':' {
				params[part[1:]] = searchParts[idx]
			}
			if part[0] == '*' && len(part) > 1 {
				params[part[1:]] = strings.Join(searchParts[idx:], "/")
				break
			}
		}
		return n, params
	}

	return nil, nil
}

// handler 拿到指定的处理函数进行处理
func (r *router) handle(ctx *Context) {
	// ctx.Method GET/POST...
	// ctx.Path /hello/xxx/xxx
	n, params := r.getRoute(ctx.Method, ctx.Path)
	if n != nil {
		ctx.Params = params
		key := fmt.Sprintf("%s-%s", ctx.Method, n.pattern)

		// 将当前handler添加到 ctx.handlers 的末尾
		ctx.handlers = append(ctx.handlers, r.handlers[key])
	} else {
		ctx.handlers = append(ctx.handlers, func(c *Context) {
			c.String(http.StatusNotFound, "404 NOT FOUND: %s", c.Path)
		})
	}

	ctx.Next()
}

// parsePattern ...
func parsePattern(pattern string) (parts []string) {
	vs := strings.Split(pattern, "/")

	for _, part := range vs {
		if part != "" {
			parts = append(parts, part)

			// 匹配任意字符
			if part[0] == '*' {
				break
			}
		}
	}

	return parts
}
