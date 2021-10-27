package gee

import (
	"fmt"
	"net/http"
)

type Router struct {
	trie     *Trie
	handlers map[string]map[string]HandleFunc
}

func NewRouter() *Router {
	return &Router{
		trie:     newTrieNode("/"),
		handlers: make(map[string]map[string]HandleFunc),
	}
}

func (router *Router) addHandler(method, pattern string, handler HandleFunc) {
	if _, ok := router.handlers[pattern]; !ok {
		router.handlers[pattern] = make(map[string]HandleFunc)
	}
	router.handlers[pattern][method] = handler
	router.trie.insert(pattern)
}

func (router *Router) handle(context *Context) {
	method := context.Method
	url := context.Path
	node, params := router.trie.search(url)
	context.Params = params
	if node == nil {
		context.Error(http.StatusNotFound, fmt.Sprintf("404 NOT FOUND: %s", url))
		return
	}
	pattern := node.pattern
	if _, ok := router.handlers[pattern][method]; !ok {
		context.Error(http.StatusMethodNotAllowed, fmt.Sprintf("405 METHOD NOT ALLOWED: %s", url))
		return
	}
	handler := router.handlers[pattern][method]
	handler(context)
}
