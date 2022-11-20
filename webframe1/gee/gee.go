package gee

import (
	"fmt"
	"log"
	"net/http"
)

type HandleFunc func(w http.ResponseWriter, r *http.Request)

type Engine struct {
	routes map[string]HandleFunc
}

func New() *Engine {
	return &Engine{
		routes: make(map[string]HandleFunc),
	}
}
func (e *Engine) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	key := r.Method + "-" + r.URL.Path
	if handle, ok := e.routes[key]; ok {
		handle(w, r)
	} else {
		log.Println(fmt.Fprintf(w, "404 not found %s", r.URL))
	}
}

func (e *Engine) addRoute(method string, pattern string, handle HandleFunc) {
	key := method + "-" + pattern
	e.routes[key] = handle
}

func (e *Engine) Get(pattern string, handle HandleFunc) {
	e.addRoute("GET", pattern, handle)
}

func (e *Engine) Post(pattern string, handle HandleFunc) {
	e.addRoute("POST", pattern, handle)
}

func (e *Engine) Delete(pattern string, handle HandleFunc) {
	e.addRoute("DELETE", pattern, handle)
}

func (e *Engine) Update(pattern string, handle HandleFunc) {
	e.addRoute("UPDATE", pattern, handle)
}

func (e *Engine) Run(addr string) error {
	return http.ListenAndServe(addr, e)
}
