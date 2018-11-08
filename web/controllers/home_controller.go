package controllers

import (
	"gbs/gream/web"
)

func init() {
	web.Register(&HomeController{})
}

type HomeController struct {
	web.BaseController
}

func (c *HomeController) IndexJson() {
	c.RenderJson(&web.H{"name": c.Param("name")})
}

func (c *HomeController) Index() {
	c.RenderText("hello " + c.Param("name"))
}
