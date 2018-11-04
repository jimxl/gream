package controller

import (
	"net/http"

	"gbs/gream/web"

	"github.com/gorilla/mux"
	"github.com/json-iterator/go"
)

type BaseController struct {
	responseWriter    http.ResponseWriter
	request           *http.Request
	RenderDefaultFile bool
	response          Response
	parameters        map[string]string
}

var jsonBuilder = jsoniter.ConfigCompatibleWithStandardLibrary

func (self *BaseController) InitFromContext(w http.ResponseWriter, r *http.Request) {
	self.responseWriter = w
	self.request = r
	self.response.StatusCode = http.StatusOK
	self.RenderDefaultFile = true
	self.parameters = mux.Vars(r)
}

func (self *BaseController) RenderText(content string) {
	self.RenderDefaultFile = false
	self.responseWriter.WriteHeader(self.response.StatusCode)
	self.responseWriter.Write([]byte(content))
}

func (self *BaseController) RenderJson(json *web.H) {
	self.RenderDefaultFile = false
	self.responseWriter.Header().Set("Content-Type", "application/json")
	content, _ := jsonBuilder.Marshal(json)
	self.responseWriter.WriteHeader(self.response.StatusCode)
	self.responseWriter.Write([]byte(content))
}

func (self *BaseController) Render() {

}

func (self *BaseController) Param(name string) string {
	return self.parameters[name]
}
