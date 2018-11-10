package admin

import (
	"gbs/gream/web/controller"
)

type HomeController struct {
	controller.BaseController
}

func (c *HomeController) Index() {
	c.RenderText("scope home controller: hello " + c.Param("name"))
}
