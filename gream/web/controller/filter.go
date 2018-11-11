package controller

import "gbs/gream/web/http_router"

type Filter interface {
	doAction(*http_router.Context) bool
}

var beforeActions = []Filter{}

func BeforeAction(filter Filter) {
	beforeActions = append(beforeActions, filter)
}
