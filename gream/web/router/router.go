package router

import (
	"net/http"
	"path/filepath"

	"gbs/gream/web/http_router"
)

var r = &router{}
var GET = r.GET

// func Scope(path string) *GScope {
// 	return scope.Scope(path)
// }

// func Namespace(path string) *GScope {
// 	return scope.Namespace(path)
// }

type router struct {
	urlSpace    string
	moduleSpace string
}

func (r *router) GET(path string, opt H) {
	route := route{
		path:        path,
		opt:         opt,
		method:      http.MethodGet,
		urlSpace:    r.urlSpace,
		moduleSpace: r.moduleSpace,
	}
	fullPath := filepath.Join(r.urlSpace, path)
	http_router.GET(fullPath, route.getHandle())
}
