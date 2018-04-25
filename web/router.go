package web

import (
	r "github.com/jimxl/gream/web/router"
)

type Router struct {
	r.Helper
}

func (self *Router) Draw() {
	self.GET("/home_json/:name", &r.Opt{To: "home#index_json"})
	// self.GET("/home_text/:name", &controllers.HomeController{}, "IndexText")
}
