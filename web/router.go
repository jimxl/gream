package web

import (
	"github.com/jimxl/gream/web/router"
	"gbs/web/controllers"
)

type Router struct {
	router.Helper
}

func (self *Router) Draw() {
	self.Get("/home_json/:name", &controllers.HomeController{}, "Index1Json")
	self.Get("/home_text/:name", &controllers.HomeController{}, "IndexText")
}