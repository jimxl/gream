package controller

import (
	"gbs/gream/web/http_router"
)

type WebController struct {
	actions map[string]ActionI
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
		actions: make(map[string]ActionI),
	}
	return wc
}

func Controller(name string, controllerBody func()) bool {
	wc, ok := controllers[name]
	if !ok {
		wc = makeWebController()
		controllers[name] = wc
	}
	currentController = wc
	controllerBody()
	return true
}

func Action(name string, actionBody ActionI) {
	currentController.addAction(name, actionBody)
}

func DoAction(cName, aName string, ctx http_router.Context) {
	if wc, ok := controllers[cName]; ok {
		if action, ok := wc.actions[aName]; ok {
			action(ctx)
		}
	}
}
