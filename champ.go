package champ

import (
	"net/http"
	"strings"
)

type router struct {
	routes map[string][]*http.HandlerFunc
}

func newRouter() *router {
	return &router{
		routes: make(map[string][]*http.HandlerFunc),
	}
}

type Handler func(w http.ResponseWriter, r *http.Request)

func (cr *router) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	handlers := cr.match(r.Method, r.URL.Path)
	for _, handler := range handlers {
		handler.ServeHTTP(w, r)
	}
}

func (cr *router) Route(method, path string, handlerFunc http.HandlerFunc) {
	cr.routes[strings.ToLower(method+path)] = append(cr.routes[strings.ToLower(method+path)], &handlerFunc)
}

func (cr *router) match(method, path string) []*http.HandlerFunc {
	return cr.routes[strings.ToLower(method+path)]
}

type champ struct {
	*router
}

func New() *champ {
	return &champ{
		router: newRouter(),
	}
}

func (c champ) ListenAndServe(port string) error {
	return http.ListenAndServe(port, c.router)
}
func (c *champ) AddMiddleware() {
	panic("implement me")
}

func (c *champ) Group(prefix string) *group {
	return &group{
		prefix: prefix,
		champ:  c,
	}
}

type group struct {
	prefix string
	*champ
}
