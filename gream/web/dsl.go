package web

import (
	"gbs/gream/web/controller"
	"gbs/gream/web/http_router"
)

type H = http_router.H
type Context = http_router.Context

var Controller = controller.Controller
var Action = controller.Action
