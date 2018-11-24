package http_router

import (
	"github.com/jimxl/gream/config"
	"github.com/kataras/iris"
)

var app = &Application{iris.Default()}

func init() {
	//app.Use(recover.New())
}

func Router() *Application {
	return app
}

func Run() {
	app.Run(iris.Addr(config.App.Addr))
}

type Application struct {
	*iris.Application
}

type Context = iris.Context

type H = iris.Map
type Handler = iris.Handler

var Handle = app.Handle
var HandleMany = app.HandleMany
var Any = app.Any
