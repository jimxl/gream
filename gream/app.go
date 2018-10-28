package gream

import (
	"gbs/gream/web"
	"gbs/gream/web/router"
)

func Run() {
	web.Debug()
	router.Run()
}
