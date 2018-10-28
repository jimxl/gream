package router

import (
	urlPath "path"

	"gbs/gream/web"
	"gbs/rgo/rarray"
	"gbs/rgo/rstring"

	"github.com/gin-gonic/gin"
)

type Scope struct {
	basePath  string
	scopePath string
}

var httpMethods = map[string](func(relativePath string, handlers ...gin.HandlerFunc) gin.IRoutes){
	"GET":     re.router.GET,
	"POST":    re.router.POST,
	"PUT":     re.router.PUT,
	"DELETE":  re.router.DELETE,
	"PATCH":   re.router.PATCH,
	"HEAD":    re.router.HEAD,
	"OPTIONS": re.router.OPTIONS,
	"Any":     re.router.Any,
}

var resourcesActions = []string{"index", "new", "create", "show", "edit", "update", "destroy"}

func (scope *Scope) GET(path string, controllerAndAction string, args ...Opt) {
	scope.handle("GET", path, controllerAndAction, args...)
}

func (scope *Scope) POST(path string, controllerAndAction string, args ...Opt) {
	scope.handle("POST", path, controllerAndAction, args...)
}

func (scope *Scope) PUT(path string, controllerAndAction string, args ...Opt) {
	scope.handle("PUT", path, controllerAndAction, args...)
}

func (scope *Scope) DELETE(path string, controllerAndAction string, args ...Opt) {
	scope.handle("DELETE", path, controllerAndAction, args...)
}

func (scope *Scope) PATCH(path string, controllerAndAction string, args ...Opt) {
	scope.handle("PATCH", path, controllerAndAction, args...)
}

func (scope *Scope) HEAD(path string, controllerAndAction string, args ...Opt) {
	scope.handle("HEAD", path, controllerAndAction, args...)
}

func (scope *Scope) OPTIONS(path string, controllerAndAction string, args ...Opt) {
	scope.handle("OPTIONS", path, controllerAndAction, args...)
}

func (scope *Scope) Any(path string, controllerAndAction string, args ...Opt) {
	scope.handle("Any", path, controllerAndAction, args...)
}

func (scope *Scope) handle(method string, path string, controllerAndAction string, args ...Opt) {
	if ginHandle, ok := httpMethods[method]; ok {
		fullPath := urlPath.Join(scope.basePath, path)
		controllerClassName, actionName, handle := requestHandle(scope.scopePath, controllerAndAction, args...)
		web.AddRInfo(&web.RInfo{
			Verb:                method,
			URI:                 fullPath,
			ControllerAndAction: controllerClassName + "#" + actionName,
		})
		ginHandle(urlPath.Join(scope.basePath, path), handle)
	}
}

func (scope *Scope) Scope(path string) *Scope {
	return &Scope{
		basePath:  urlPath.Join(scope.basePath, path),
		scopePath: urlPath.Join(scope.basePath, path),
	}
}

func (scope *Scope) Namespace(path string) *Scope {
	return &Scope{
		basePath: urlPath.Join(scope.basePath, path),
	}
}

// Resources TODO: 要实现Only和Expect方法，用于关闭某些默认路由
// TODO:  由于gin不支持 GET /users/new  和 GET /users/:id 这样的路由，会冲突只能换成 GET /user/:id
func (scope *Scope) Resources(name string, opts ...Opt) *Scope {
	var actions map[string]bool
	if len(opts) == 1 {
		opt := opts[0]
		actions = rarray.OnlyAndExpect(opt.Only, opt.Expect, resourcesActions)
	}

	pluralName := rstring.Plural(name)
	singularName := rstring.Singular(name)

	if len(actions) == 0 || actions["index"] {
		scope.GET("/"+pluralName, pluralName+"#index")
	}
	if len(actions) == 0 || actions["new"] {
		scope.GET("/"+singularName+"/new", pluralName+"#new")
	}
	if len(actions) == 0 || actions["create"] {
		scope.POST("/"+name, pluralName+"#create")
	}
	if len(actions) == 0 || actions["show"] {
		scope.GET("/"+pluralName+"/:id", pluralName+"#show")
	}
	if len(actions) == 0 || actions["edit"] {
		scope.GET("/"+singularName+"/edit", pluralName+"#edit")
	}
	if len(actions) == 0 || actions["update"] {
		scope.PATCH("/"+pluralName+"/:id", pluralName+"#update")
		scope.PUT("/"+pluralName+"/:id", pluralName+"#update")
	}
	if len(actions) == 0 || actions["destroy"] {
		scope.DELETE("/"+pluralName+"/:id", pluralName+"#destroy")
	}
	return &Scope{
		basePath: urlPath.Join(scope.basePath, name),
	}
}
