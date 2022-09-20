参考项目：https://geektutu.com/post/gee-day2.html

启动项目：
~~~go
func main() {
	r := gee.Default()
	r.GET("/", func(ctx *gee.Context) {
		ctx.JSON(http.StatusOK, gee.H{
			"status": 1, "msg": "hello world",
		})
	})
	r.Run()
}
