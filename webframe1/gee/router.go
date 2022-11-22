package gee

import (
	"log"
	"net/http"
	"strings"
)

type HandleFunc func(ctx *Context)

type router struct {
	handlers map[string]HandleFunc
	root     map[string]*node
}

func newRouter() *router {
	return &router{
		handlers: make(map[string]HandleFunc),
		root:     make(map[string]*node),
	}
}

func parsePattern(pattern string) []string {
	splits := strings.Split(pattern, "/")
	parts := make([]string, 0)
	for _, part := range splits {
		if part != "" {
			parts = append(parts, part)
			if part[0] == '*' {
				break
			}
		}
	}
	return parts
}

func (r *router) addRouter(method string, path string, handler HandleFunc) {

	parts := parsePattern(path)
	root, ok := r.root[method]
	if !ok {
		r.root[method] = &node{}
		root = r.root[method]
	}
	root.insert(path, parts, -1)
	key := method + "-" + path
	r.handlers[key] = handler
	log.Println("[gee]\t", method, "\t", path)
}

func (r *router) getRouter(method string, path string) (*node, map[string]string) {
	if _, ok := r.root[method]; !ok {
		return nil, nil
	}
	searchParts := parsePattern(path)
	routerNode := r.root[method].search(searchParts, -1)
	if routerNode == nil {
		return nil, nil
	}
	parts := parsePattern(routerNode.pattern)
	params := make(map[string]string)
	for index, part := range parts {
		if part[0] == ':' {
			params[part[1:]] = searchParts[index]
		}
		if part[0] == '*' {
			params[part[1:]] = strings.Join(searchParts[index:], "/")
		}
	}
	return routerNode, params
}

func (r *router) handle(context *Context) {
	node, params := r.getRouter(context.Method, context.Path)
	if node != nil {
		key := context.Method + "-" + node.pattern
		context.Params = params
		r.handlers[key](context)
	} else {
		context.String(http.StatusNotFound, "404 not Found")
	}
}
