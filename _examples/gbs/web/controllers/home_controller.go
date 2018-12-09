package controllers

import (
	"fmt"
	. "github.com/jimxl/gream/web"
)

func init() {
	Controller("home")

	BeforeAll(allFilter)
	BeforeAction("index", indexFilter)

	Action("index", func(ctx Context) {
		ctx.SetSession("key", "value1")
		ctx.RenderTextf("hello, %s, %s", ctx.Params("name"), ctx.Session("key"))
	})

	Action("test_render", func(ctx Context) {
		ctx.ToView("test", "test123")
	})
}

func indexFilter(ctx Context) bool {
	fmt.Printf("index filter, controller %s, action: %s\n", ctx.ControllerName, ctx.ActionName)
	return true
}

func allFilter(ctx Context) bool {
	fmt.Printf("all filter, controller %s, action: %s\n", ctx.ControllerName, ctx.ActionName)
	return true
}
