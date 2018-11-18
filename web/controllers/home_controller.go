package controllers

import (
	. "gbs/gream/web"
)

func init() {
	Controller("home")

	Action("index", func(ctx Context) {
		ctx.Writef("hello, %s", ctx.Params().GetString("name"))
	})
}
