package controller

import "gbs/gream/web/http_router"

type Controller interface {
	InitFromContext(*http_router.Context)
	RenderText(string, ...interface{})
	RenderJson(*http_router.H)
	Render()
	Param(string) string
}
