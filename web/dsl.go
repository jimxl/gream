package web

import (
	"github.com/jimxl/gream/web/controller"
	"github.com/jimxl/gream/web/http_router"
)

type H = http_router.H

//type Context = http_router.Context

type Context = controller.Context

var Controller = controller.Controller
var Action = controller.Action

var BeforeAction = controller.BeforeAction
var AfterAction = controller.AfterAction
var BeforeAll = controller.BeforeAll
var AfterAll = controller.AfterAll
