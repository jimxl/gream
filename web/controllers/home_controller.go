package controllers

import (
	"fmt"
	"gbs/gream/web"
)

type HomeController struct {
	web.BaseController
}

func (c *HomeController) Init() {
	fmt.Println("init controller")
	c.BeforeAction("Index", func() bool {
		fmt.Println("index action before action")
		return true
	})
}

func (c *HomeController) IndexJson() {
	c.RenderJson(&web.H{"name": c.Param("name")})
}

func (c *HomeController) Index() {
	c.RenderText("hello " + c.Param("name"))
}
