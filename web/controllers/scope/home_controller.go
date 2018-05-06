package scope

import (
	"github.com/jimxl/gream/web"
	"github.com/jimxl/gream/web/controller"
)

func init() {
	web.Register(&HomeController{})
}

type HomeController struct {
	controller.BaseController
}

func (c *HomeController) Index() {
	c.RenderText("scope home controller: hello " + c.Param("name"))
}
