package controllers

import (
	"github.com/jimxl/gream/web/controller"
	"github.com/jimxl/gream"
)

type HomeController struct {
	controller.BaseController
}

func(self *HomeController) IndexJsonAction () {
	self.RenderJson(&gream.H{"name": self.Param("name")})
}

func(self *HomeController) IndexTextAction () {
	self.RenderText("hello " + self.Param("name"))
}
