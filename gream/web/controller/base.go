package controller

import (
	"errors"
	"fmt"
	"gbs/gream/logger"
	"net/http"
	"reflect"

	"gbs/gream/web/http_router"
)

type BaseController struct {
	FilterModule
	context           *http_router.Context
	renderDefaultFile bool
}

func (controller *BaseController) InitFromContext_(c *http_router.Context) {
	controller.context = c
	controller.renderDefaultFile = true
}

func (controller *BaseController) RenderText(format string, values ...interface{}) {
	controller.renderDefaultFile = false
	controller.context.String(http.StatusOK, format)
}

func (controller *BaseController) RenderJson(json *http_router.H) {
	controller.renderDefaultFile = false
	controller.context.JSON(http.StatusOK, json)
}

func (controller *BaseController) renderView() {

}

func (controller *BaseController) Param(name string) string {
	return controller.context.Param(name)
}

func (controller *BaseController) Init() {
	fmt.Println("===============")
}

func (controller *BaseController) CallAction_(actionName string, self reflect.Value) {
	action := self.MethodByName(actionName)
	// TODO: 为啥golang当中指针一样，却不能反射呢
	//fmt.Printf("3 %+v\n", self.Elem().Type())
	//fmt.Printf("4 %+v\n", reflect.ValueOf(controller).Elem().Type())
	//fmt.Printf("5 %+v\n", controller)
	//fmt.Printf("6 %+v\n", self.Pointer())
	//fmt.Printf("7 %+v\n", reflect.ValueOf(controller).Pointer())
	if !action.IsValid() {
		err := errors.New("action invalid")
		logger.Error(err.Error())
		return
	}
	action.Call([]reflect.Value{})

	if controller.renderDefaultFile {
		controller.renderView()
	}
}
