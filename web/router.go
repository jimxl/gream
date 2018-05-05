package web

import (
	r "github.com/jimxl/gream/web/router"
)

type Router struct {
	r.Helper
}

func (self *Router) Draw() {
	self.GET("/home_json/:name", "home#index_json")
}
