package main

import (
	"gee"
	"net/http"
)

func main() {
	r := gee.Default()
	r.GET("/", func(ctx *gee.Context) {
		ctx.JSON(http.StatusOK, gee.H{
			"status": 1, "msg": "hello world",
		})
	})
	r.Run()
}
