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

func (c *HomeController) IndexJson() {
	c.RenderJson(&web.H{"name": c.Param("name")})
}

func (c *HomeController) IndexText() {
	c.RenderText("hello " + c.Param("name"))
}
