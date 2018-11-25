package router

import (
	"github.com/jimxl/gream/web/controller"
	"github.com/jimxl/gream/web/http_router"
	"github.com/kataras/iris"
	"path/filepath"
	"regexp"

	"github.com/jimxl/gream/rgo/rstring"
)

var controllerScopeRe = regexp.MustCompile("(\\w*/)?(\\w*)#(\\w*)$")

type route struct {
	path        string
	opt         H
	controller  string
	action      string
	urlSpace    string
	moduleSpace string
	fullpath    string
}

func (s *route) getHandle() func(iris.Context) {
	s.parseControllerAndAction()
	return func(ctx iris.Context) {
		controller.DoAction_(s.controller, s.action, http_router.NewContext(ctx))
	}
}

func (s *route) parseControllerAndAction() {
	info := controllerScopeRe.FindStringSubmatch(rstring.Downcase(s.opt["to"]))
	controllerName, action, dir := rstring.Downcase(info[2]), rstring.Downcase(info[3]), rstring.Downcase(info[1])
	s.controller = filepath.Join(s.moduleSpace, dir, controllerName)
	s.action = action
}
