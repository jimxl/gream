package middleware

import (
	"github.com/jimxl/gream/web/http_router"
)

var app *http_router.Application

func init() {
	app = http_router.Router()
	//http_router.Use(loggerMiddleWare())
}
