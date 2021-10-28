package gee

type RouterGroup struct {
	prefix   string
	parent   *RouterGroup
	children []*RouterGroup
	engine   *Engine
}

func (group *RouterGroup) Group(prefix string) *RouterGroup {
	newGroup := new(RouterGroup)
	*newGroup = RouterGroup{
		prefix:   group.prefix + prefix,
		parent:   group,
		children: make([]*RouterGroup, 0),
		engine:   group.engine,
	}
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
