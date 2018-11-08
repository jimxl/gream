package controller

import (
	"net/http"

	"gbs/gream/web/http_router"
)

type BaseController struct {
	context *http_router.Context

	RenderDefaultFile bool
}

func (controller *BaseController) InitFromContext(c *http_router.Context) {
	controller.context = c
	controller.RenderDefaultFile = true
}

func (controller *BaseController) RenderText(format string, values ...interface{}) {
	controller.RenderDefaultFile = false
	controller.context.String(http.StatusOK, format)
}

func (controller *BaseController) RenderJson(json *http_router.H) {
	controller.RenderDefaultFile = false
	controller.context.JSON(http.StatusOK, json)
}

func (controller *BaseController) Render() {

}

func (controller *BaseController) Param(name string) string {
	return controller.context.Param(name)
}
