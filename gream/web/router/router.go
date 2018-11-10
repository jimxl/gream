package router

import (
	"gbs/rgo/rmap"
	"gbs/rgo/rstring"
	"net/http"
	"path/filepath"

	"gbs/gream/web/http_router"
)

var ApplicationRouterDraw func()

var r = &Router{}

var (
	HEAD      = r.HEAD
	GET       = r.GET
	POST      = r.POST
	PUT       = r.PUT
	PATCH     = r.PATCH
	DELETE    = r.DELETE

	Resources = r.Resources
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

func (r *Router) POST(path string, opt H) {
	route := r.getRoute(http.MethodPost, path, opt)
	http_router.POST(route.fullpath, route.getHandle())
}

func (r *Router) HEAD(path string, opt H) {
	route := r.getRoute(http.MethodPost, path, opt)
	http_router.HEAD(route.fullpath, route.getHandle())
}

func (r *Router) PUT(path string, opt H) {
	route := r.getRoute(http.MethodPost, path, opt)
	http_router.PUT(route.fullpath, route.getHandle())
}

func (r *Router) PATCH(path string, opt H) {
	route := r.getRoute(http.MethodPost, path, opt)
	http_router.PATCH(route.fullpath, route.getHandle())
}

func (r *Router) DELETE(path string, opt H) {
	route := r.getRoute(http.MethodDelete, path, opt)
	http_router.DELETE(route.fullpath, route.getHandle())
}

func (r *Router) Any(path string, opt H) {
	route := r.getRoute(http.MethodDelete, path, opt)
	http_router.Any(route.fullpath, route.getHandle())
}

func (r *Router) Resources(name string, opts ...H) {
	pluralName, singularName := rstring.Plural(name), rstring.Singular(name)

	var opt H
	if len(opts) <= 0 {
		opt = H{}
	} else {
		opt = opts[0]
	}

	indexOpt := H{"to": pluralName + "#index"}
	rmap.Merge(&indexOpt, opt)
	r.GET(urlJoin("/", pluralName), indexOpt)

	newOpt := H{"to": pluralName + "#new"}
	rmap.Merge(&newOpt, opt)
	r.GET(urlJoin("/", singularName, "new"), newOpt)

	createOpt := H{"to": pluralName + "#create"}
	rmap.Merge(&newOpt, opt)
	r.POST(urlJoin("/", pluralName), createOpt)

	showOpt := H{"to": pluralName + "#show"}
	rmap.Merge(&showOpt, opt)
	r.GET(urlJoin("/", pluralName, ":id"), showOpt)

	editOpt := H{"to": pluralName + "#edit"}
	rmap.Merge(&editOpt, opt)
	r.GET(urlJoin("/", pluralName, ":id/edit"), editOpt)

	updateOpt := H{"to": pluralName + "#update"}
	rmap.Merge(&updateOpt, opt)
	r.PATCH(urlJoin("/", pluralName, ":id"), updateOpt)
	r.PUT(urlJoin("/", pluralName, ":id"), updateOpt)

	deleteOpt := H{"to": pluralName + "#destroy"}
	rmap.Merge(&updateOpt, opt)
	r.DELETE(urlJoin("/", pluralName, ":id"), deleteOpt)
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
	}

	if optModule, ok := opt["module"]; ok {
		route.moduleSpace = urlJoin(optModule, r.moduleSpace)
	}
	route.fullpath = urlJoin(route.urlSpace, path)
	return &route
}
