package main

import (
	"webframe1/gee"
)

func main() {
	r := gee.New()
	r.Get("/user/:id", func(ctx *gee.Context) {
		ctx.String(200, "成功获取 id 为%s的user", ctx.Param("id"))
	})
	r.Get("/user/class/*path", func(ctx *gee.Context) {
		ctx.String(200, ctx.Param("path"))
	})
	r.Run("localhost:9999")
}
