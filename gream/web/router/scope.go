package router

import (
	"errors"
	"fmt"
	"net/http"
	"reflect"
	"regexp"

	"gbs/gream/logger"
	"gbs/gream/web"
	. "gbs/gream/web/response"
	"gbs/rgo/rstring"

	"github.com/gorilla/mux"
)

type Scope struct {
	route *mux.Route
	path  string
}

func (scope *Scope) handle(controllerAndAction string) *mux.Route {
	controllerName, actionName, dir := getName(controllerAndAction)
	controllerClassName := scope.path + "/" + dir + controllerName + "Controller"
	return scope.route.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// TODO: 要加上类型转换判断,防止错误
		response := w.(*Response)
		controller, err := createController(controllerClassName, response, r)
		if err != nil {
			w.WriteHeader(http.StatusNotFound)
			w.Write([]byte(err.Error()))
			return
		}

		fmt.Println(actionName)
		err = callAction(controller, actionName, controllerName)
		if err != nil {
			w.WriteHeader(http.StatusNotFound)
			w.Write([]byte(err.Error()))
			return
		}
	})
}

var controllerScopeRe = regexp.MustCompile("(.*/)?(.*)#(.*)$")

func getName(controllerAndAction string) (controller, action, dir string) {
	info := controllerScopeRe.FindStringSubmatch(rstring.Downcase(controllerAndAction))
	controller = rstring.Capitalize(info[2])
	action = rstring.Capitalize(info[3])
	dir = info[1]
	return
}

func createController(name string, response *Response, r *http.Request) (*reflect.Value, error) {
	controllerType := web.GetController(name)
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
	method.Call([]reflect.Value{reflect.ValueOf(response), reflect.ValueOf(r)})
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
