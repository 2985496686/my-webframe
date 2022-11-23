package gee

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type Context struct {
	//Object data
	Writer  http.ResponseWriter
	Request *http.Request
	//Request data
	Method string
	Path   string
	//Response data
	StatusCode int
	//路径中的参数
	Params map[string]string
	//middlewares
	handles []HandleFunc
	index   int
}

func newContext(w http.ResponseWriter, r *http.Request) *Context {
	return &Context{
		Writer:     w,
		Request:    r,
		Method:     r.Method,
		Path:       r.URL.Path,
		StatusCode: 400,
		index:      -1,
	}
}

func (ctx *Context) Next() {
	ctx.index++
	for ; ctx.index < len(ctx.handles); ctx.index++ {
		ctx.handles[ctx.index](ctx)
	}
}

func (ctx *Context) PostForm(key string) string {
	return ctx.Request.FormValue(key)
}

func (ctx *Context) Query(key string) string {
	return ctx.Request.URL.Query().Get(key)
}

func (ctx *Context) Status(code int) {
	ctx.StatusCode = code
	ctx.Writer.WriteHeader(code)
}

func (ctx *Context) SetHeader(key, value string) {
	ctx.Writer.Header().Set(key, value)
}

func (ctx *Context) JSON(code int, obj interface{}) {
	ctx.Status(code)
	ctx.SetHeader("Content-Type", "application/json")
	encoder := json.NewEncoder(ctx.Writer)
	if err := encoder.Encode(obj); err != nil {
		http.Error(ctx.Writer, err.Error(), http.StatusInternalServerError)
	}
}

func (ctx *Context) String(code int, format string, a ...any) {
	ctx.Status(code)
	ctx.SetHeader("Content-Type", "text/plain")
	ctx.Writer.Write([]byte(fmt.Sprintf(format, a...)))
}

func (ctx *Context) HTML(code int, html string) {
	ctx.Status(code)
	ctx.SetHeader("Content-Type", "text/html")
	ctx.Writer.Write([]byte(html))
}

func (ctx *Context) Param(key string) string {
	return ctx.Params[key]
}

func (ctx *Context) Fail(code int, errMessage string) {
	ctx.Status(code)
	ctx.Writer.Write([]byte(errMessage))
}
