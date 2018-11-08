package web

import (
	"gbs/gream/web/controller"
	"gbs/gream/web/http_router"
)

type BaseController = controller.BaseController

type Controller = controller.Controller

type H = http_router.H

func Register(c Controller) {
	controller.Register(c)
}
