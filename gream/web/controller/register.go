package controller

import (
	"reflect"
	"regexp"

	"gbs/gream/logger"
)

var controllers = map[string]reflect.Type{}
var controllerScopeRe = regexp.MustCompile("web/controllers?(.*)$")

func Register(controller Controller) {
	controllerType := reflect.TypeOf(controller)
	controllerScope := controllerScopeRe.FindStringSubmatch(controllerType.Elem().PkgPath())[1]
	controllers[controllerScope+"/"+controllerType.Elem().Name()] = controllerType
}

func GetController(name string) reflect.Type {
	return controllers[name]
}

func PrintControllers() {
	for name, controller := range controllers {
		logger.Debugf("%v => %v", name, controller)
	}
}
