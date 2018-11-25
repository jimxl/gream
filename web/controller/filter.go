package controller

//
import (
	"github.com/jimxl/gream/rgo/rstring"
	"github.com/jimxl/gream/web/http_router"
)

// TODO: 对于All filter 可以把action的名字传进去, 方便判断。还可以实现skip_filter 这样的方法
// TODO: 还能需要实现only, except方法

type Filter = func(*http_router.Context) bool

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

// TODO: 先判断Application级别的filter是否存在, 如果存在要先制性, 另外可能需要用Hash的方式来存储filter, 这样当skip_action的时候可以直接用名字代替
func (wc *WebController) CallActionWithFilter(actionName string, ctx *http_router.Context) {
	name := rstring.Downcase(actionName)
	ctx.ActionName = actionName
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
