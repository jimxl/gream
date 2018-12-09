package controllers

import (
	. "github.com/jimxl/gream/web"
)

func init() {
	Controller("users")

	Action("index", func(ctx Context) {
		ctx.RenderText("users controller index action")
	})

	Action("new", func(ctx Context) {
		ctx.RenderText("users controller new action")
	})

	Action("create", func(ctx Context) {
		ctx.RenderText("users controller create action")
	})

	Action("show", func(ctx Context) {
		ctx.RenderText("users controller show action")
	})

	Action("edit", func(ctx Context) {
		ctx.RenderText("users controller edit action")
	})

	Action("update", func(ctx Context) {
		ctx.RenderText("users controller update action")
	})

	Action("destroy", func(ctx Context) {
		ctx.RenderText("users controller destroy action")
	})
}
