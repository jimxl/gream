package gream

import (
	_ "gbs/gream/config"
	_ "gbs/gream/logger"
	_ "gbs/gream/web/middleware"

	"gbs/gream/web"
)

func Run() {
	web.Debug()
	web.Run()
}
