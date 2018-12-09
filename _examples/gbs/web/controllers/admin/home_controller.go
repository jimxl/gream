package admin

import . "github.com/jimxl/gream/web"

func init() {
	Controller("admin/home")

	Action("index", func(ctx Context) {
		ctx.RenderTextf("admin home controller hello, %s, %s", ctx.ControllerName(), ctx.ActionName())
	})
}
