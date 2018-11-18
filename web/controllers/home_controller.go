package controllers

import (
	. "gbs/gream/web/controller"
	"gbs/gream/web/http_router"
)

var _ = Controller("home", func() {
	i := 1
	Action("index", func(ctx http_router.Context) {
		i = i + 1
		ctx.Writef("%s, %v", "hello", i)
	})
})
