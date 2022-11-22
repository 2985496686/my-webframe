package main

import (
	"time"
	"webframe1/gee"
	"webframe1/middlewares"
)

func main() {
	r := gee.New()
	r1 := r.Group("/user")
	{
		r1.Use(middlewares.Log)
		r1.Get("/:id", func(ctx *gee.Context) {
			time.Sleep(1000)
			ctx.String(200, "成功获取 id 为%s的user", ctx.Param("id"))
		})
		r1.Get("/class/*path", func(ctx *gee.Context) {
			ctx.String(200, ctx.Param("path"))
		})
	}

	r2 := r.Group("/admin")
	r2.Get("/add", func(ctx *gee.Context) {
		ctx.String(200, "admin")
	})

	r.Run("localhost:9999")
}
