package controllers

import (
	"fmt"
	. "gbs/gream/web"
)

func init() {
	Controller("home")

	BeforeAll(allFilter)
	BeforeAction("index", indexFilter)

	Action("index", func(ctx Context) {
		ctx.Writef("hello, %s", ctx.Params().GetString("name"))
	})
}

func indexFilter(ctx Context) bool {
	fmt.Printf("index filter, controller %s, action: %s\n", ctx.Values().Get("controller"), ctx.Values().Get("action"))
	return true
}

func allFilter(ctx Context) bool {
	fmt.Printf("all filter, controller %s, action: %s\n", ctx.Values().Get("controller"), ctx.Values().Get("action"))
	return true
}
