package controllers

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

func (c *HomeController) IndexJsonAction() {
	c.RenderJson(&web.H{"name": c.Param("name")})
}

func (c *HomeController) IndexTextAction() {
	c.RenderText("hello " + c.Param("name"))
}
