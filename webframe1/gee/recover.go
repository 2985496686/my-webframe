package gee

import (
	"fmt"
	"log"
	"net/http"
	"runtime"
	"strings"
)

func trace(message string) string {
	var pcs [32]uintptr
	n := runtime.Callers(3, pcs[:]) // skip first 3 caller

	var str strings.Builder
	str.WriteString(message + "\nTraceback:")
	for _, pc := range pcs[:n] {
		fn := runtime.FuncForPC(pc)
		file, line := fn.FileLine(pc)
		str.WriteString(fmt.Sprintf("\n\t%s:%d", file, line))
	}
	return str.String()
}

func Recover() HandleFunc {
	return func(ctx *Context) {
		defer func() {
			if err := recover(); err != nil {
				log.Println(trace(err.(string)))
				ctx.Fail(http.StatusInternalServerError, "Internal server error")
			}
		}()
		ctx.Next()
	}
}
