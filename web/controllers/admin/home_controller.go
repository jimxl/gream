package admin

import . "gbs/gream/web"

func init() {
	Controller("admin/home")

	Action("index", func(ctx Context) {
		ctx.Writef("admin home controller hello, %s", ctx.Params().GetString("name"))
	})
}
