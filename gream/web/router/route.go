package router

import (
	"errors"
	"fmt"
	"net/http"
	"path/filepath"
	"reflect"
	"regexp"

	"gbs/gream/logger"
	"gbs/gream/web/controller"
	"gbs/gream/web/http_router"
	"gbs/rgo/rstring"
)

var controllerScopeRe = regexp.MustCompile("(\\w*/)?(\\w*)#(\\w*)$")

type route struct {
	path        string
	method      string
	opt         H
	controller  string
	action      string
	urlSpace    string
	moduleSpace string

	controllerInstance *reflect.Value
}

func (s *route) getHandle() func(*http_router.Context) {
	s.parseControllerAndAction()
	return func(c *http_router.Context) {
		err := s.createController(c)
		if err != nil {
			c.String(http.StatusNotFound, err.Error())
			return
		}

		err = s.callAction()
		if err != nil {
			c.String(http.StatusNotFound, err.Error())
			return
		}
	}
}

func (s *route) parseControllerAndAction() {
	info := controllerScopeRe.FindStringSubmatch(rstring.Downcase(s.opt["to"]))
	controller, action, dir := rstring.Capitalize(info[2]), rstring.Capitalize(info[3]), info[1]
	s.controller = filepath.Join("/", s.moduleSpace, dir, controller) + "Controller"
	s.action = action
}

func (s *route) createController(c *http_router.Context) error {
	fmt.Println(s)
	controllerType := controller.GetController(s.controller)
	if controllerType == nil {
		err := errors.New("controller invalid")
		logger.Error(err.Error())
		return err
	}
	controllerInstance := reflect.New(controllerType.Elem())
	method := controllerInstance.MethodByName("InitFromContext")
	if !method.IsValid() {
		err := errors.New("controller invalid")
		logger.Error(err.Error())
		return err
	}
	method.Call([]reflect.Value{reflect.ValueOf(c)})
	s.controllerInstance = &controllerInstance
	return nil
}

func (s *route) callAction() error {
	action := s.controllerInstance.MethodByName(s.action)
	if !action.IsValid() {
		err := errors.New("action invalid")
		logger.Error(err.Error())
		return err
	}
	action.Call([]reflect.Value{})

	// if controller.FieldByName("RenderDefaultFile").Bool() {
	s.controllerInstance.MethodByName("Render").Call([]reflect.Value{})
	// }
	return nil
}
