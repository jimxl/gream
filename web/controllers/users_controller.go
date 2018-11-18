package controllers

import (
	. "gbs/gream/web"
)

func init() {
	Controller("users")

	Action("index", func(ctx Context) {
		ctx.WriteString("users controller index action")
	})

	Action("new", func(ctx Context) {
		ctx.WriteString("users controller new action")
	})

	Action("create", func(ctx Context) {
		ctx.WriteString("users controller create action")
	})

	Action("show", func(ctx Context) {
		ctx.WriteString("users controller show action")
	})

	Action("edit", func(ctx Context) {
		ctx.WriteString("users controller edit action")
	})

	Action("update", func(ctx Context) {
		ctx.WriteString("users controller update action")
	})

	Action("destroy", func(ctx Context) {
		ctx.WriteString("users controller destroy action")
	})
}
