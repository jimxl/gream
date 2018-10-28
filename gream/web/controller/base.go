package controller

import (
	"net/http"

	"gbs/gream/web"

	"github.com/gin-gonic/gin"
)

type BaseController struct {
	context *gin.Context

	RenderDefaultFile bool
	response          Response
}

func (self *BaseController) InitFromGinContext(context *gin.Context) {
	self.response.StatusCode = http.StatusOK
	self.RenderDefaultFile = true
	self.context = context
}

func (self *BaseController) RenderText(content string) {
	self.RenderDefaultFile = false
	self.context.String(self.response.StatusCode, content)
}

func (self *BaseController) RenderJson(json *web.H) {
	self.RenderDefaultFile = false
	self.context.JSON(self.response.StatusCode, json)
}

func (self *BaseController) Render() {

}

func (self *BaseController) Param(name string) string {
	return self.context.Param(name)
}

func (self *BaseController) Query(name string) string {
	return self.context.Query(name)
}

func (self *BaseController) DefaultQuery(name string, defaultValue string) string {
	return self.context.DefaultQuery(name, defaultValue)
}

func (self *BaseController) PostForm(name string) string {
	return self.context.PostForm(name)
}

func (self *BaseController) DefaultPostForm(name string, defaultValue string) string {
	return self.context.DefaultPostForm(name, defaultValue)
}
