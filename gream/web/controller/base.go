package controller

import (
	"net/http"

	"gbs/gream/web"
	. "gbs/gream/web/response"

	"github.com/gorilla/mux"
	"github.com/json-iterator/go"
)

type BaseController struct {
	request           *http.Request
	RenderDefaultFile bool
	response          *Response
	parameters        map[string]string
}

var jsonBuilder = jsoniter.ConfigCompatibleWithStandardLibrary

func (self *BaseController) InitFromContext(response *Response, r *http.Request) {
	self.request = r
	self.response = response
	self.response.WriteHeader(http.StatusOK)
	self.RenderDefaultFile = true
	self.parameters = mux.Vars(r)
}

func (self *BaseController) RenderText(content string) {
	self.RenderDefaultFile = false
	self.response.Write([]byte(content))
}

func (self *BaseController) RenderJson(json *web.H) {
	self.RenderDefaultFile = false
	self.response.Header().Set("Content-Type", "application/json")
	content, _ := jsonBuilder.Marshal(json)
	self.response.Write([]byte(content))
}

func (self *BaseController) Render() {

}

func (self *BaseController) Param(name string) string {
	return self.parameters[name]
}
