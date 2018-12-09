package view

import (
	"github.com/gobuffalo/plush"
	"github.com/jimxl/gream/config"
	"github.com/jimxl/gream/logger"
	"io/ioutil"
)

type View struct {
	tmplFile    string
	tmplContext *plush.Context
}

func Create(tmplFile string) *View {
	return &View{
		tmplFile:    config.App.Path("web", "views", tmplFile),
		tmplContext: plush.NewContext(),
	}
}

func (view *View) Set(name string, value interface{}) {
	view.tmplContext.Set(name, value)
}

func (view *View) Render() string {
	s, err := plush.Render(view.html(), view.tmplContext)
	if err != nil {
		logger.Error(err)
		return ""
	}
	return s
}

func (view *View) html() string {
	// TODO: 生产模式下可以cache住
	c, err := ioutil.ReadFile(view.tmplFile)
	if err != nil {
		logger.Error(err)
		return ""
	}
	return string(c)
}
