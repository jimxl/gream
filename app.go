package gream

import (
	_ "github.com/jimxl/gream/config"
	_ "github.com/jimxl/gream/logger"
	_ "github.com/jimxl/gream/web/middleware"

	"github.com/jimxl/gream/web"
)

func Run() {
	web.Debug()
	web.Run()
}
