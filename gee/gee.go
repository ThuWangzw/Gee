package gee

import (
	"net/http"
	"strings"
)

type HandleFunc func(*Context)

type Engine struct {
	*RouterGroup
	routers *Router
	groups  []*RouterGroup
}

func New() *Engine {
	engine := new(Engine)
	*engine = Engine{
		routers: NewRouter(),
		RouterGroup: &RouterGroup{
			children:    make([]*RouterGroup, 0),
			engine:      engine,
			middleWares: make([]HandleFunc, 0),
		},
		groups: make([]*RouterGroup, 0),
	}
	engine.groups = append(engine.groups, engine.RouterGroup)
	return engine
}

func (gee *Engine) ServeHTTP(w http.ResponseWriter, req *http.Request) {
	context := NewContext(w, req)
	// find group
	for _, group := range gee.groups {
		if strings.HasPrefix(context.Path, group.prefix) {
			context.handlers = append(context.handlers, group.middleWares...)
		}
	}
	context.handlerIdx = -1
	gee.routers.handle(context)
}

func (gee *Engine) Run(addr string) {
	http.ListenAndServe(addr, gee)
}
