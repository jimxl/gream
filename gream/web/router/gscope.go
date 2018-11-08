package router

import (
	"errors"
	"net/http"
	"path/filepath"
	"reflect"
	"regexp"

	"gbs/gream/logger"
	"gbs/gream/web/controller"
	"gbs/gream/web/http_router"
	"gbs/rgo/rstring"
)

type GScope struct {
	path           string
	classNamespace string
}

// func (s *GScope) Scope(name string) *GScope {
// 	s.route.PathPrefix("/" + name)
// 	s.path = filepath.Join(s.path, name)
// 	return s
// }

// func (s *GScope) Namespace(name string) *GScope {
// 	s.route.PathPrefix("/" + name)
// 	return s
// }

func (s *GScope) GET(path, controllerAndAction string) {
	fullPath := filepath.Join(s.path, path)
	http_router.GET(fullPath, s.handle(controllerAndAction))
}

func (s *GScope) handle(controllerAndAction string) func(*http_router.Context) {
	controllerName, actionName, dir := getName(controllerAndAction)
	controllerClassName := filepath.Join("/", s.path, dir, controllerName) + "Controller"
	return func(c *http_router.Context) {
		logger.Debugf("controller: %s, action: %s in %s", controllerName, actionName, controllerClassName)
		controller, err := createController(controllerClassName, c)
		if err != nil {
			c.String(http.StatusNotFound, err.Error())
			return
		}

		err = callAction(controller, actionName, controllerName)
		if err != nil {
			c.String(http.StatusNotFound, err.Error())
			return
		}
	}
}

var controllerScopeRe = regexp.MustCompile("(\\w*/)?(\\w*)#(\\w*)$")

func getName(controllerAndAction string) (controller, action, dir string) {
	info := controllerScopeRe.FindStringSubmatch(rstring.Downcase(controllerAndAction))
	controller = rstring.Capitalize(info[2])
	action = rstring.Capitalize(info[3])
	dir = info[1]
	return
}

func createController(name string, c *http_router.Context) (*reflect.Value, error) {
	controllerType := controller.GetController(name)
	if controllerType == nil {
		err := errors.New("controller invalid")
		logger.Error(err.Error())
		return nil, err
	}
	controllerInstance := reflect.New(controllerType.Elem())
	method := controllerInstance.MethodByName("InitFromContext")
	if !method.IsValid() {
		err := errors.New("controller invalid")
		logger.Error(err.Error())
		return nil, err
	}
	method.Call([]reflect.Value{reflect.ValueOf(c)})
	return &controllerInstance, nil
}

func callAction(controller *reflect.Value, name string, controllerName string) error {
	action := controller.MethodByName(name)
	if !action.IsValid() {
		err := errors.New("action invalid")
		logger.Error(err.Error())
		return err
	}
	action.Call([]reflect.Value{})

	// if controller.FieldByName("RenderDefaultFile").Bool() {
	controller.MethodByName("Render").Call([]reflect.Value{})
	// }
	return nil
}

// // Resources TODO: 要实现Only和Expect方法，用于关闭某些默认路由
// // TODO:  由于gin不支持 GET /users/new  和 GET /users/:id 这样的路由，会冲突只能换成 GET /user/:id
// func (scope *Scope) Resources(name string, opts ...Opt) *Scope {
// 	var actions map[string]bool
// 	if len(opts) == 1 {
// 		opt := opts[0]
// 		actions = rarray.OnlyAndExpect(opt.Only, opt.Expect, resourcesActions)
// 	}

// 	pluralName := rstring.Plural(name)
// 	singularName := rstring.Singular(name)

// 	if len(actions) == 0 || actions["index"] {
// 		scope.GET("/"+pluralName, pluralName+"#index")
// 	}
// 	if len(actions) == 0 || actions["new"] {
// 		scope.GET("/"+singularName+"/new", pluralName+"#new")
// 	}
// 	if len(actions) == 0 || actions["create"] {
// 		scope.POST("/"+name, pluralName+"#create")
// 	}
// 	if len(actions) == 0 || actions["show"] {
// 		scope.GET("/"+pluralName+"/:id", pluralName+"#show")
// 	}
// 	if len(actions) == 0 || actions["edit"] {
// 		scope.GET("/"+singularName+"/edit", pluralName+"#edit")
// 	}
// 	if len(actions) == 0 || actions["update"] {
// 		scope.PATCH("/"+pluralName+"/:id", pluralName+"#update")
// 		scope.PUT("/"+pluralName+"/:id", pluralName+"#update")
// 	}
// 	if len(actions) == 0 || actions["destroy"] {
// 		scope.DELETE("/"+pluralName+"/:id", pluralName+"#destroy")
// 	}
// 	return &Scope{
// 		basePath: urlPath.Join(scope.basePath, name),
// 	}
// }
