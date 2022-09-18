package main

import (
	"gee"
	"net/http"
)

func main() {
	r := gee.New()
	r.GET("/", func(ctx *gee.Context) {
		ctx.JSON(http.StatusOK, gee.H{
			"status": 1,
			"msg":    "welcome",
		})
	})
	r.Run(":8080")
}
