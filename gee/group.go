package gee

type RouterGroup struct {
	prefix      string
	parent      *RouterGroup
	children    []*RouterGroup
	middleWares []HandleFunc
	engine      *Engine
}

func (group *RouterGroup) Group(prefix string) *RouterGroup {
	newGroup := new(RouterGroup)
	*newGroup = RouterGroup{
		prefix:      group.prefix + prefix,
		parent:      group,
		children:    make([]*RouterGroup, 0),
		engine:      group.engine,
		middleWares: make([]HandleFunc, 0),
	}
	newGroup.engine.groups = append(newGroup.engine.groups, newGroup)
	group.children = append(group.children, newGroup)
	return newGroup
}

func (group *RouterGroup) Get(pattern string, handler HandleFunc) {
	pattern = group.prefix + pattern
	group.engine.routers.addHandler("GET", pattern, handler)
}

func (group *RouterGroup) Post(pattern string, handler HandleFunc) {
	pattern = group.prefix + pattern
	group.engine.routers.addHandler("POST", pattern, handler)
}

func (group *RouterGroup) Use(middlewares ...HandleFunc) {
	group.middleWares = append(group.middleWares, middlewares...)
}
