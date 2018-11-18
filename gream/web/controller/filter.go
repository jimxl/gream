package controller

//
import (
	"gbs/gream/web/http_router"
	"gbs/rgo/rstring"
)

// TODO: 对于All filter 可以把action的名字传进去, 方便判断。还可以实现skip_filter 这样的方法
// TODO: 还能需要实现only, except方法

type Filter = func(http_router.Context) bool

func (wc *WebController) BeforeAll(filter Filter) {
	wc.beforeAll = append(wc.beforeAll, filter)
}

func (wc *WebController) AfterAll(filter Filter) {
	wc.afterAll = append(wc.afterAll, filter)
}

func (wc *WebController) BeforeAction(actionName string, filter Filter) {
	name := rstring.Downcase(actionName)
	if _, ok := wc.beforeActions[name]; !ok {
		wc.beforeActions[name] = []Filter{}
	}
	wc.beforeActions[name] = append(wc.beforeActions[name], filter)
}

func (wc *WebController) AfterAction(actionName string, filter Filter) {
	name := rstring.Downcase(actionName)
	if _, ok := wc.afterActions[name]; !ok {
		wc.afterActions[name] = []Filter{}
	}
	wc.afterActions[name] = append(wc.afterActions[name], filter)
}

func (wc *WebController) CallActionWithFilter(actionName string, ctx http_router.Context) {
	name := rstring.Downcase(actionName)
	ctx.Values().Set("action", name)
	if action, ok := wc.actions[name]; ok {

		for _, filter := range wc.beforeAll {
			if !filter(ctx) {
				return
			}
		}

		for _, filter := range wc.beforeActions[name] {
			if !filter(ctx) {
				return
			}
		}

		action(ctx)

		for _, filter := range wc.afterActions[name] {
			if !filter(ctx) {
				return
			}
		}

		for _, filter := range wc.afterAll {
			if !filter(ctx) {
				return
			}
		}
	}
}
