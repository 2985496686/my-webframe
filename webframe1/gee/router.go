package gee

import (
	"log"
	"net/http"
)

type HandleFunc func(ctx *Context)

type router struct {
	handlers map[string]HandleFunc
}

func newRouter() router {
	return router{
		handlers: make(map[string]HandleFunc),
	}
}

func (r *router) addRouter(method string, path string, handler HandleFunc) {
	key := method + "-" + path
	log.Printf("[gee debug] %4s \t %s ", method, path)
	r.handlers[key] = handler
}

func (r *router) handle(context *Context) {
	key := context.Method + "-" + context.Path
	if handle, ok := r.handlers[key]; ok {
		handle(context)
	} else {
		context.String(http.StatusNotFound, "404 not Found")
	}
}
