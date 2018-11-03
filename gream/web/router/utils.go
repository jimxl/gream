package router

import (
	"errors"
	"net/http"
	"reflect"
	"regexp"

	"gbs/rgo/rstring"

	"gbs/gream/logger"
	"gbs/gream/web"

	"github.com/gin-gonic/gin"
)

func Run() {
	re.run()
}

type Opt struct {
	To     string
	As     string
	Only   []string
	Expect []string
}

var controllerScopeRe = regexp.MustCompile("(.*/)?(.*)#(.*)$")

// 这个地方可以方便的再加入另外的opt参数，方便做参数的定制化 具体可以参考rails的方式
func requestHandle(path string, controllerAndAction string) (http.ResponseWriter, *http.Request) {
	controllerName, actionName, path := getName(controllerAndAction)
	controllerClassName := scopePath + "/" + path + controllerName + "Controller"
	return controllerClassName, actionName, func(c *gin.Context) {
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

func getName(controllerAndAction string) (controller, action, path string) {
	info := controllerScopeRe.FindStringSubmatch(rstring.Downcase(controllerAndAction))
	controller = rstring.Capitalize(info[2])
	action = rstring.Capitalize(info[3])
	path = info[1]
	return
}

func createController(name string, c *gin.Context) (*reflect.Value, error) {
	controllerType := web.GetController(name)
	if controllerType == nil {
		err := errors.New("controller invalid")
		logger.Error(err.Error())
		return nil, err
	}
	controllerInstance := reflect.New(controllerType.Elem())
	method := controllerInstance.MethodByName("InitFromGinContext")
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
