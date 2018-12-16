package controller

import (
	"github.com/jimxl/gream/web/http_router"
)

type WebController struct {
	parent  string
	actions map[string]ActionI

	beforeActions map[string][]filterN
	afterActions  map[string][]filterN

	beforeAll []filterN
	afterAll  []filterN

	skippedBeforeAll    []string
	skippedBeforeAction map[string]string
}

func (wc *WebController) addAction(name string, actionBody ActionI) {
	wc.actions[name] = actionBody
}

var controllers map[string]*WebController
var currentController *WebController

func init() {
	controllers = make(map[string]*WebController)
}

func makeWebController() *WebController {
	wc := &WebController{
		actions:             make(map[string]ActionI),
		beforeActions:       make(map[string][]filterN),
		afterActions:        make(map[string][]filterN),
		beforeAll:           []filterN{},
		afterAll:            []filterN{},
		skippedBeforeAll:    []string{},
		skippedBeforeAction: make(map[string]string),
	}
	return wc
}

func Controller(name string, parent ...string) bool {
	wc, ok := controllers[name]
	if !ok {
		wc = makeWebController()
		controllers[name] = wc
	}
	currentController = wc
	if len(parent) >= 1 {
		wc.parent = parent[0]
	} else {
		wc.parent = "application"
	}
	return true
}

func Action(name string, actionBody ActionI) {
	currentController.addAction(name, actionBody)
}

func BeforeAction(filter Filter, opts ...H) {
	currentController.BeforeAction(filter, opts...)
}

func AfterAction(filter Filter, opts ...H) {
	currentController.AfterAction(filter, opts...)
}

func DoAction_(cName, aName string, ctx *http_router.Context) {
	if wc, ok := controllers[cName]; ok {
		ctx.ControllerName_ = cName
		wc.callActionWithFilter(aName, ctx)
		ctx.Render()
	}
}
