package middlewares

import (
	"log"
	"time"
	"webframe1/gee"
)

func Log(ctx *gee.Context) {
	start := time.Now().UnixNano()
	ctx.Next()
	log.Println("耗时:", time.Now().UnixNano()-start)
}
