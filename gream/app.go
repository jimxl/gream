package gream

import (
	_ "gbs/gream/config"
	_ "gbs/gream/logger"
	_ "gbs/gream/web/middleware"

	"gbs/gream/web"
	"gbs/gream/web/router"
)

func Run() {
	web.Debug()
	router.Run()
}
