package gee

import (
	"fmt"
	"net/http"
)

type HandleFunc func(http.ResponseWriter, *http.Request)

type Engine struct {
	routers map[string]map[string]HandleFunc
}

func New() *Engine {
	routers := make(map[string]map[string]HandleFunc)
	return &Engine{
		routers: routers,
	}
}

func (gee *Engine) addHandler(method, url string, handler HandleFunc) {
	if _, ok := gee.routers[url]; !ok {
		gee.routers[url] = make(map[string]HandleFunc)
	}
	gee.routers[url][method] = handler
}

func (gee *Engine) Get(url string, handler HandleFunc) {
	gee.addHandler("GET", url, handler)
}

func (gee *Engine) Post(url string, handler HandleFunc) {
	gee.addHandler("POST", url, handler)
}

func (gee *Engine) ServeHTTP(w http.ResponseWriter, req *http.Request) {
	method := req.Method
	url := req.URL.Path
	if _, ok := gee.routers[url]; !ok {
		w.WriteHeader(http.StatusNotFound)
		fmt.Fprintf(w, "404 NOT FOUND: %s", url)
		return
	}
	if _, ok := gee.routers[url][method]; !ok {
		w.WriteHeader(http.StatusMethodNotAllowed)
		fmt.Fprintf(w, "405 METHOD NOT ALLOWED: %s: %s", method, url)
		return
	}
	handler := gee.routers[url][method]
	handler(w, req)
}

func (gee *Engine) Run(addr string) {
	http.ListenAndServe(addr, gee)
}
