package http_router

import (
	"gbs/gream/config"
	"github.com/kataras/iris"
)

var app *Application

func init() {
	app = &Application{iris.Default()}
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
type HandlerFunc = func(c Context)

func Use(_ HandlerFunc) {
	//	#TODO: 加上中间件的使用
}

func GET(path string, handler HandlerFunc) {
	app.Get(path, func(c Context) {
		handler(c)
	})
}

func POST(path string, handler HandlerFunc) {
	app.Post(path, func(c Context) {
		handler(c)
	})
}

func DELETE(path string, handler HandlerFunc) {
	app.Delete(path, func(c Context) {
		handler(c)
	})
}

func PATCH(path string, handler HandlerFunc) {
	app.Patch(path, func(c Context) {
		handler(c)
	})
}

func PUT(path string, handler HandlerFunc) {
	app.Put(path, func(c Context) {
		handler(c)
	})
}

func OPTIONS(path string, handler HandlerFunc) {
	app.Options(path, func(c Context) {
		handler(c)
	})
}

func HEAD(path string, handler HandlerFunc) {
	app.Head(path, func(c Context) {
		handler(c)
	})
}

func Any(path string, handler HandlerFunc) {
	app.Any(path, func(c Context) {
		handler(c)
	})
}
