package web

import (
	"gbs/gream/web/http_router"
	"gbs/gream/web/router"
)

func Run() {
	http_router.Run()
}

func Debug() {
	// PrintControllers()
	router.PrintUrls()
}
