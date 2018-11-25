package http_router

import (
	"github.com/jimxl/gream/logger"
	"github.com/kataras/iris"
	"path/filepath"
)

func NewContext(ctx iris.Context) *Context {
	return &Context{
		irisContext: ctx,
		isRender:    false,
	}
}

type Context struct {
	irisContext iris.Context

	ControllerName string
	ActionName     string

	isRender bool
}

func (ctx *Context) Params(name string) string {
	return ctx.irisContext.Params().Get(name)
}

func (ctx *Context) SetSession(key, value string) {
	// TODO: add session options, 例如 有效时间等
	ctx.irisContext.SetCookieKV(key, value)
}

func (ctx *Context) Session(key string) string {
	return ctx.irisContext.GetCookie(key)
}

func (ctx *Context) RenderText(body string) error {
	_, err := ctx.irisContext.WriteString(body)
	ctx.isRender = true
	return err
}

func (ctx *Context) RenderTextf(format string, values ...interface{}) error {
	_, err := ctx.irisContext.Writef(format, values...)
	ctx.isRender = true
	return err
}

func (ctx *Context) Render(name ...string) {
	if !ctx.isRender {
		ctx.isRender = true
		var templateFilePath string
		// TODO: 获取的view文件的最后可以加上请求的类型 例如, xxx.html 目前都是默认html
		if len(name) <= 0 {
			templateFilePath = filepath.Join(ctx.ControllerName, ctx.ActionName+".html")
		} else {
			templateFilePath = filepath.Join(name[0] + ".html")
		}
		logger.Info("渲染html view:" + templateFilePath)
	}
}

func (ctx *Context) RedirectTo(url string, statusHeader ...int) {
	ctx.irisContext.Redirect(url, statusHeader...)
}
