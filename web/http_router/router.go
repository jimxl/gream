package http_router

import (
	"github.com/gorilla/securecookie"
	"github.com/jimxl/gream/config"
	"github.com/kataras/iris"
	"github.com/kataras/iris/sessions"
)

var (
	app          = &Application{iris.Default()}
	secureCookie = securecookie.New([]byte(config.App.SessionHashKey), []byte(config.App.SessionBlockKey))
	sess         = sessions.New(sessions.Config{
		Cookie:       config.App.SessionID,
		Encode:       secureCookie.Encode,
		Decode:       secureCookie.Decode,
		AllowReclaim: true,
	})
)

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

type H = iris.Map
type Handler = iris.Handler

var Handle = app.Handle
var HandleMany = app.HandleMany
var Any = app.Any
