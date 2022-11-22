package gee

import (
	"log"
	"net/http"
)

type RouterGroup struct {
	prefix      string
	middlewares []HandleFunc
	parent      *RouterGroup
	engine      *Engine
}

type Engine struct {
	*router
	*RouterGroup
	routerGroups []*RouterGroup
}

func New() *Engine {
	engine := &Engine{router: newRouter()}
	engine.RouterGroup = &RouterGroup{engine: engine}
	engine.routerGroups = []*RouterGroup{engine.RouterGroup}
	return engine
}
func (e *Engine) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	context := newContext(w, r)
	e.handle(context)
	log.Printf("[gin] |  %d  |\t\t%s\t%s", context.StatusCode, context.Method, context.Path)
}

func (r *RouterGroup) Group(prefix string) *RouterGroup {
	routerGroup := &RouterGroup{
		prefix: r.prefix + prefix,
		parent: r,
		engine: r.engine,
	}
	r.engine.routerGroups = append(r.engine.routerGroups, routerGroup)
	return routerGroup
}

func (r *RouterGroup) addRoute(method string, pattern string, handle HandleFunc) {
	pattern = r.prefix + pattern
	r.engine.addRouter(method, pattern, handle)
}

func (r *RouterGroup) Get(pattern string, handle HandleFunc) {
	r.addRoute("GET", pattern, handle)
}

func (r *RouterGroup) Post(pattern string, handle HandleFunc) {
	r.addRoute("POST", pattern, handle)
}

func (r *RouterGroup) Delete(pattern string, handle HandleFunc) {
	r.addRoute("DELETE", pattern, handle)
}

func (r *RouterGroup) Update(pattern string, handle HandleFunc) {
	r.addRoute("UPDATE", pattern, handle)
}

func (e *Engine) Run(addr string) error {
	return http.ListenAndServe(addr, e)
}
