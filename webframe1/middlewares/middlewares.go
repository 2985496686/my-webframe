package middlewares

import (
	"fmt"
	"log"
	"time"
	"webframe1/gee"
)

func Log(ctx *gee.Context) {
	start := time.Now().UnixNano()
	ctx.Next()
	log.Println("耗时:", time.Now().UnixNano()-start)
}

func Hello(ctx *gee.Context) {
	fmt.Println("hello 1")
	ctx.Next()
	fmt.Println("hello 2")
}
