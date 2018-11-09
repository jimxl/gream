package router

import (
	"net/http"
	"path/filepath"

	"gbs/gream/web/http_router"
)

var r = &Router{}

var (
	GET       = r.GET
	Namespace = r.Namespace
	Scope     = r.Scope
)

func urlJoin(paths ...string) string {
	return filepath.Join(paths...)
}

type Router struct {
	urlSpace    string
	moduleSpace string
}

func (r *Router) GET(path string, opt H) {
	route := r.getRoute(http.MethodGet, path, opt)
	http_router.GET(route.fullpath, route.getHandle())
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

func (r *Router) getRoute(method, path string, opt H) *route {
	route := route{
		path:        path,
		opt:         opt,
		method:      http.MethodGet,
		urlSpace:    r.urlSpace,
		moduleSpace: r.moduleSpace,
	}
	if optPath, ok := opt["path"]; ok {
		route.urlSpace = urlJoin(optPath, r.urlSpace)
		delete(opt, "path")
	}

	if optModule, ok := opt["module"]; ok {
		route.moduleSpace = urlJoin(optModule, r.moduleSpace)
		delete(opt, "module")
	}
	route.fullpath = urlJoin(route.urlSpace, path)
	return &route
}
