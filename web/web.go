package web

import (
	"github.com/jimxl/gream/web/http_router"
	"github.com/jimxl/gream/web/router"
)

func Run() {
	router.ApplicationRouterDraw()
	http_router.Run()
}

func Debug() {
	// PrintControllers()
	router.PrintUrls()
}
