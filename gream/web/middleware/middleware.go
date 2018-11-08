package middleware

import (
	"gbs/gream/web/http_router"
)

var re *http_router.Engine

func init() {
	re = http_router.Router()
	http_router.Use(loggerMiddleWare())
}
