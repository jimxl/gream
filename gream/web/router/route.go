package router

import (
	"gbs/gream/web/controller"
	"path/filepath"
	"regexp"

	"gbs/gream/web/http_router"
	"gbs/rgo/rstring"
)

var controllerScopeRe = regexp.MustCompile("(\\w*/)?(\\w*)#(\\w*)$")

type route struct {
	path        string
	method      string
	opt         H
	controller  string
	action      string
	urlSpace    string
	moduleSpace string
	fullpath    string
}

func (s *route) getHandle() func(http_router.Context) {
	s.parseControllerAndAction()
	return func(ctx http_router.Context) {
		controller.DoAction_(s.controller, s.action, ctx)
	}
}

func (s *route) parseControllerAndAction() {
	info := controllerScopeRe.FindStringSubmatch(rstring.Downcase(s.opt["to"]))
	controllerName, action, dir := rstring.Downcase(info[2]), rstring.Downcase(info[3]), rstring.Downcase(info[1])
	s.controller = filepath.Join(s.moduleSpace, dir, controllerName)
	s.action = action
}
