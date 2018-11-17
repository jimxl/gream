package controller

import (
	"errors"
	"fmt"
	"gbs/gream/logger"
	"reflect"
)

type Filter = func() bool

var beforeActions = []Filter{}

func BeforeAction(filter Filter) {
	beforeActions = append(beforeActions, filter)
}

func (controller *BaseController) CallActionWithFilter(actionName string, self reflect.Value) {
	for _, filter := range controller.beforeActions[actionName] {
		if !filter() {
			return
		}
	}
	controller.callAction(actionName, self)
	for _, filter := range controller.afterActions[actionName] {
		if !filter() {
			return
		}
	}
}

func (controller *BaseController) callAction(actionName string, self reflect.Value) {
	action := self.MethodByName(actionName)
	fmt.Printf("3 %+v\n", self.Elem().Type())
	fmt.Printf("4 %+v\n", reflect.ValueOf(controller).Elem().Type())
	fmt.Printf("5 %+v\n", controller)

	fmt.Printf("6 %+v\n", self.Pointer())
	fmt.Printf("7 %+v\n", reflect.ValueOf(controller).Pointer())
	if !action.IsValid() {
		err := errors.New("action invalid1")
		logger.Error(err.Error())
		return
	}
	action.Call([]reflect.Value{})

	// if controller.FieldByName("RenderDefaultFile").Bool() {
	self.MethodByName("Render").Call([]reflect.Value{})
	// }
}
