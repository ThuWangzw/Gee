package gee

import (
	"net/http"
)

type HandleFunc func(*Context)

type Engine struct {
	routers *Router
}

func New() *Engine {
	return &Engine{
		routers: NewRouter(),
	}
}

func (gee *Engine) Get(url string, handler HandleFunc) {
	gee.routers.addHandler("GET", url, handler)
}

func (gee *Engine) Post(url string, handler HandleFunc) {
	gee.routers.addHandler("POST", url, handler)
}

func (gee *Engine) ServeHTTP(w http.ResponseWriter, req *http.Request) {
	context := NewContext(w, req)
	gee.routers.handle(context)
}

func (gee *Engine) Run(addr string) {
	http.ListenAndServe(addr, gee)
}
