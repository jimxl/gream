package web

import (
	"github.com/jimxl/gream/web/controller"
)

//type H = http_router.H
type H = map[string]string

//type Context = http_router.Context

type Context = controller.Context

var Controller = controller.Controller
var Action = controller.Action

var BeforeAction = controller.BeforeAction
var AfterAction = controller.AfterAction
