package gee

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type H map[string]interface{}

type Context struct {
	W   http.ResponseWriter
	Req *http.Request

	Path   string
	Method string
	Params map[string]string

	ResponceStatus int

	handlers   []HandleFunc
	handlerIdx int
}

func NewContext(w http.ResponseWriter, req *http.Request) *Context {
	return &Context{
		W:          w,
		Req:        req,
		Path:       req.URL.Path,
		Method:     req.Method,
		handlers:   make([]HandleFunc, 0),
		handlerIdx: -1,
	}
}

func (context *Context) Next() {
	for context.handlerIdx++; context.handlerIdx < len(context.handlers); context.handlerIdx++ {
		context.handlers[context.handlerIdx](context)
	}
}

func (context *Context) Param(key string) string {
	return context.Params[key]
}

func (context *Context) PostForm(name string) string {
	return context.Req.FormValue(name)
}

func (context *Context) Query(name string) string {
	return context.Req.URL.Query().Get(name)
}

func (context *Context) Status(StatusCode int) {
	context.ResponceStatus = StatusCode
	context.W.WriteHeader(StatusCode)
}

func (context *Context) SetHeader(key, value string) {
	context.W.Header().Set(key, value)
}

func (context *Context) Error(code int, msg string) {
	http.Error(context.W, msg, code)
}

func (context *Context) JSON(code int, obj interface{}) {
	context.SetHeader("Content-Type", "application/json")
	context.Status(code)
	encoder := json.NewEncoder(context.W)
	if err := encoder.Encode(obj); err != nil {
		context.Error(http.StatusInternalServerError, err.Error())
	}
}

func (context *Context) String(code int, format string, a ...interface{}) {
	context.SetHeader("Content-Type", "text/plain")
	context.Status(code)
	if _, err := fmt.Fprintf(context.W, format, a...); err != nil {
		context.Error(http.StatusInternalServerError, err.Error())
	}
}

func (context *Context) HTML(code int, html string) {
	context.SetHeader("Content-Type", "text/html")
	context.Status(code)
	if _, err := fmt.Fprint(context.W, html); err != nil {
		context.Error(http.StatusInternalServerError, err.Error())
	}
}
