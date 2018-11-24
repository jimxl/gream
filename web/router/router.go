package router

import (
	"github.com/jimxl/gream/rgo/rmap"
	"net/http"
	"path/filepath"

	"github.com/jimxl/gream/web/http_router"
)

var ApplicationRouterDraw func()

var r = &Router{}

var (
	HEAD   = r.HEAD
	GET    = r.GET
	POST   = r.POST
	PUT    = r.PUT
	PATCH  = r.PATCH
	DELETE = r.DELETE

	Resources = r.Resources
	Namespace = r.Namespace
	Scope     = r.Scope
)

func urlJoin(paths ...string) string {
	return filepath.Join(paths...)
}

func withOpt(to string, opts []H, f func(opt H)) {
	if len(opts) <= 0 {
		f(H{"to": to})
	} else {
		opt := H{"to": to}
		rmap.Merge(&opt, opts[0])
		f(opt)
	}
}

func handle(method, path, to string, opts []H) {
	withOpt(to, opts, func(opt H) {
		route := r.getRoute(path, opt)
		http_router.Handle(method, route.fullpath, route.getHandle())
	})
}

type Router struct {
	urlSpace    string
	moduleSpace string
}

func (r *Router) GET(path, to string, opts ...H) {
	handle(http.MethodGet, path, to, opts)
}

func (r *Router) POST(path, to string, opts ...H) {
	handle(http.MethodPost, path, to, opts)
}

func (r *Router) HEAD(path, to string, opts ...H) {
	handle(http.MethodHead, path, to, opts)
}

func (r *Router) PUT(path, to string, opts ...H) {
	handle(http.MethodPut, path, to, opts)
}

func (r *Router) PATCH(path, to string, opts ...H) {
	handle(http.MethodPatch, path, to, opts)
}

func (r *Router) DELETE(path, to string, opts ...H) {
	handle(http.MethodDelete, path, to, opts)
}

func (r *Router) Any(path, to string, opts ...H) {
	withOpt(to, opts, func(opt H) {
		route := r.getRoute(path, opt)
		http_router.Any(route.fullpath, route.getHandle())
	})
}

func (r *Router) Namespace(space string) *Router {
	router := Router{
		urlSpace:    urlJoin(r.urlSpace, space),
		moduleSpace: urlJoin(r.moduleSpace, space),
	}
	return &router
}

func (r *Router) Scope(arg interface{}) *Router {
	router := Router{}
	switch arg.(type) {
	case string:
		router.urlSpace = urlJoin(r.urlSpace, arg.(string))
		router.moduleSpace = r.moduleSpace
	case H:
		router.urlSpace = r.urlSpace
		moduleSpace := arg.(H)["module"]
		router.moduleSpace = urlJoin(r.moduleSpace, moduleSpace)
	}
	return &router
}

func (r *Router) getRoute(path string, opt H) *route {
	route := route{
		path:        path,
		opt:         opt,
		urlSpace:    r.urlSpace,
		moduleSpace: r.moduleSpace,
	}
	if optPath, ok := opt["path"]; ok {
		route.urlSpace = urlJoin(optPath, r.urlSpace)
	}

	if optModule, ok := opt["module"]; ok {
		route.moduleSpace = urlJoin(optModule, r.moduleSpace)
	}
	route.fullpath = urlJoin(route.urlSpace, path)
	return &route
}
