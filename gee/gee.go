package gee

import (
	"fmt"
	"log"
	"net/http"
)

func init() {
	log.SetPrefix("[gee] ")
	log.SetFlags(log.LstdFlags | log.Lshortfile)
}

type Engine struct {
	/* key -> method-pattern */
	router map[string]http.HandlerFunc
}

func New() *Engine {
	return &Engine{router: make(map[string]http.HandlerFunc)}
}

// ServeHTTP 从router中拿到指定的handler进行处理
func (e Engine) ServeHTTP(w http.ResponseWriter, req *http.Request) {
	key := fmt.Sprintf("%s-%s", req.Method, req.URL.Path)
	handler, ok := e.router[key]
	if !ok {
		log.Println("404 not found: ", req.Method)
		return
	}
	handler(w, req)
}

func (e Engine) Run(address string) {
	err := http.ListenAndServe(address, e)
	if err != nil {
		log.Fatal("start web server error: ", err)
	}
}

func (e Engine) addRouter(method, pattern string, handler http.HandlerFunc) {
	key := fmt.Sprintf("%s-%s", method, pattern)
	e.router[key] = handler
}

func (e Engine) GET(pattern string, handler http.HandlerFunc) {
	e.addRouter("GET", pattern, handler)
}

func (e Engine) POST(pattern string, handler http.HandlerFunc) {
	e.addRouter("POST", pattern, handler)
}

func (e Engine) DELETE(pattern string, handler http.HandlerFunc) {
	e.addRouter("DELETE", pattern, handler)
}

func (e Engine) PUT(pattern string, handler http.HandlerFunc) {
	e.addRouter("PUT", pattern, handler)
}

func (e Engine) PATCH(pattern string, handler http.HandlerFunc) {
	e.addRouter("PATCH", pattern, handler)
}
