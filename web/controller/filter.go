package controller

//
import (
	"github.com/jimxl/gream/logger"
	"github.com/jimxl/gream/rgo/rstring"
	"github.com/jimxl/gream/web/http_router"
	"reflect"
	"runtime"
)

// TODO: 对于All filter 可以把action的名字传进去, 方便判断。还可以实现skip_filter 这样的方法
// TODO: 还能需要实现only, except方法

type Filter = func(Context) bool
type filterN struct {
	run  Filter
	name string
}

type H = map[string]string

func makeFilterN(filter Filter) filterN {
	return filterN{filter, getFilterName(filter)}
}

func getFilterName(filter Filter) string {
	fullname := runtime.FuncForPC(reflect.ValueOf(filter).Pointer()).Name()
	isMatch, results := rstring.Match(fullname, "\\.(.*)")
	if isMatch {
		return results[0]
	}
	return ""
}

func (wc *WebController) BeforeAction(filter Filter, opts ...H) {
	if len(opts) <= 0 {
		wc.addBeforeAll(filter)
	} else {
		for _, name := range rstring.Split(opts[0]["only"], ",") {
			wc.addBeforeAction(name, filter)
		}
	}
}

func (wc *WebController) AfterAction(filter Filter, opts ...H) {
	if len(opts) <= 0 {
		wc.addAfterAll(filter)
	} else {
		for _, name := range rstring.Split(opts[0]["only"], ",") {
			wc.addAfterAction(name, filter)
		}
	}
}

func (wc *WebController) addBeforeAll(filter Filter) {
	wc.beforeAll = append(wc.beforeAll, makeFilterN(filter))
}

func (wc *WebController) addAfterAll(filter Filter) {
	wc.afterAll = append(wc.afterAll, makeFilterN(filter))
}

func (wc *WebController) addBeforeAction(actionName string, filter Filter) {
	if _, ok := wc.beforeActions[actionName]; !ok {
		wc.beforeActions[actionName] = []filterN{}
	}
	wc.beforeActions[actionName] = append(wc.beforeActions[actionName], makeFilterN(filter))
}

//func (wc *WebController) deleteBeforeAction(actionName string, filterName string) {
//	if filters, ok := wc.beforeActions[actionName]; ok {
//		for index, filter := range filters {
//			if filter.name == filterName {
//				copy(filters[index:], filters[index+1:])
//				return array[:len(array)-1]
//				wc.beforeActions[actionName] = filters[:len(filters)]
//			}
//		}
//	}
//}

func (wc *WebController) addAfterAction(actionName string, filter Filter) {
	if _, ok := wc.afterActions[actionName]; !ok {
		wc.afterActions[actionName] = []filterN{}
	}
	wc.afterActions[actionName] = append(wc.afterActions[actionName], makeFilterN(filter))
}

// TODO: 先判断Application级别的filter是否存在, 如果存在要先执行, 另外可能需要用Hash的方式来存储filter, 这样当skip_action的时候可以直接用名字代替
func (wc *WebController) callActionWithFilter(actionName string, ctx *http_router.Context) {
	parentController := controllers[wc.parent]
	if wc.parent != "application" && parentController == nil {
		logger.Error("非application controller 必须直接或间接继承 application controller")
		return //原理上不能出现没有parent的controller, 除非application controller
	}

	name := rstring.Downcase(actionName)
	ctx.ActionName_ = actionName
	if action, ok := wc.actions[name]; ok {

		for _, filter := range wc.beforeAll {
			if !filter.run(ctx) {
				return
			}
		}

		for _, filter := range wc.beforeActions[name] {
			if !filter.run(ctx) {
				return
			}
		}

		action(ctx)

		for _, filter := range wc.afterActions[name] {
			if !filter.run(ctx) {
				return
			}
		}

		for _, filter := range wc.afterAll {
			if !filter.run(ctx) {
				return
			}
		}
	}
}

func (wc *WebController) callBeforeAction(actionName string, ctx *http_router.Context) bool {
	return true
}

func (wc *WebController) callAfterAction(actionName string, ctx *http_router.Context) bool {
	return true
}
