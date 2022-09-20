package main

import (
	"gee"
	"net/http"
)

func main() {
	r := gee.Default()

	r.GET("/", func(c *gee.Context) {
		c.String(http.StatusOK, "Hello world\n")
	})

	r.GET("/panic", func(c *gee.Context) {
		names := []string{"hello"}
		// 测试Recovery中间件
		c.String(http.StatusOK, names[100])
	})

	r.Run(":8080")
}
