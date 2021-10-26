package gee

import (
	"fmt"
	"net/http"
)

type Router map[string]map[string]HandleFunc

func NewRouter() *Router {
	return &Router{}
}

func (router *Router) addHandler(method, url string, handler HandleFunc) {
	if _, ok := (*router)[url]; !ok {
		(*router)[url] = make(map[string]HandleFunc)
	}
	(*router)[url][method] = handler
}

func (router *Router) handle(context *Context) {
	method := context.Method
	url := context.Path
	if _, ok := (*router)[url]; !ok {
		context.Error(http.StatusNotFound, fmt.Sprintf("404 NOT FOUND: %s", url))
		return
	}
	if _, ok := (*router)[url][method]; !ok {
		context.Error(http.StatusNotFound, fmt.Sprintf("404 NOT FOUND: %s", url))
		return
	}
	handler := (*router)[url][method]
	handler(context)
}
