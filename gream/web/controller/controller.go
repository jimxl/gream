package controller

import "gbs/gream/web/http_router"

type Controller interface {
	InitFromContext_(*http_router.Context)
	RenderText(string, ...interface{})
	RenderJson(*http_router.H)
	Param(string) string

	Init()
}
