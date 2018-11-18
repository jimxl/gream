package controller

import (
	"gbs/gream/web/http_router"
)

type WebController struct {
	actions map[string]ActionI

	beforeActions map[string][]Filter
	afterActions  map[string][]Filter

	beforeAll []Filter
	afterAll  []Filter
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
		actions:       make(map[string]ActionI),
		beforeActions: make(map[string][]Filter),
		afterActions:  make(map[string][]Filter),
		beforeAll:     []Filter{},
		afterAll:      []Filter{},
	}
	return wc
}

func Controller(name string) bool {
	wc, ok := controllers[name]
	if !ok {
		wc = makeWebController()
		controllers[name] = wc
	}
	currentController = wc
	return true
}

func Action(name string, actionBody ActionI) {
	currentController.addAction(name, actionBody)
}

func BeforeAction(name string, filter Filter) {
	currentController.BeforeAction(name, filter)
}

func AfterAction(name string, filter Filter) {
	currentController.AfterAction(name, filter)
}

func BeforeAll(filter Filter) {
	currentController.BeforeAll(filter)
}

func AfterAll(filter Filter) {
	currentController.AfterAll(filter)
}

func DoAction_(cName, aName string, ctx http_router.Context) {
	if wc, ok := controllers[cName]; ok {
		ctx.Values().Set("controller", cName)
		wc.CallActionWithFilter(aName, ctx)
	}
}
