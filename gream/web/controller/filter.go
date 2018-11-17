package controller

import (
	"gbs/rgo/rstring"
	"reflect"
)

type Filter = func() bool

type FilterModule struct {
	beforeActions map[string][]Filter
	afterActions  map[string][]Filter
}

func (f *FilterModule) BeforeAction(actionName string, filter Filter) {
	if len(f.beforeActions) == 0 {
		f.beforeActions = make(map[string][]Filter)
	}
	name := rstring.Downcase(actionName)
	f.beforeActions[name] = append(f.beforeActions[name], filter)
}

func (f *FilterModule) afterAction(actionName string, filter Filter) {
	name := rstring.Downcase(actionName)
	f.afterActions[name] = append(f.afterActions[name], filter)
}

func (f *FilterModule) CallActionWithFilter(actionName string, self reflect.Value) {
	name := rstring.Downcase(actionName)
	for _, filter := range f.beforeActions[name] {
		if !filter() {
			return
		}
	}
	action := self.MethodByName("CallAction_")
	if action.IsValid() {
		in := []reflect.Value{reflect.ValueOf(actionName), reflect.ValueOf(self)}
		action.Call(in)
	}
	for _, filter := range f.afterActions[name] {
		if !filter() {
			return
		}
	}
}
