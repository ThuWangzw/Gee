package gee

import (
	"net/http"
)

type HandleFunc func(*Context)

type Engine struct {
	*RouterGroup
	routers *Router
}

func New() *Engine {
	engine := new(Engine)
	*engine = Engine{
		routers: NewRouter(),
		RouterGroup: &RouterGroup{
			children: make([]*RouterGroup, 0),
			engine:   engine,
		},
	}
	return engine
}

func (gee *Engine) ServeHTTP(w http.ResponseWriter, req *http.Request) {
	context := NewContext(w, req)
	gee.routers.handle(context)
}

func (gee *Engine) Run(addr string) {
	http.ListenAndServe(addr, gee)
}
