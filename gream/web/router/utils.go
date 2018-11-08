package router

import (
	"errors"
	"reflect"
	"regexp"

	"gbs/gream/logger"
	"gbs/gream/web/controller"
	"gbs/gream/web/http_router"
	"gbs/rgo/rstring"
)

type Opt struct {
	To     string
	As     string
	Only   []string
	Expect []string
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
