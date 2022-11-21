package gee

import (
	"net/http"
)

type Engine struct {
	router
}

func New() *Engine {
	return &Engine{
		router: newRouter(),
	}
}
func (e *Engine) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	context := newContext(w, r)
	e.handle(context)
}

func (e *Engine) addRoute(method string, pattern string, handle HandleFunc) {
	e.addRouter(method, pattern, handle)
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
