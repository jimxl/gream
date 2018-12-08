package http_router

import (
	"github.com/jimxl/gream/logger"
	"github.com/kataras/iris"
	"github.com/kataras/iris/sessions"
	"path/filepath"
)

func NewContext(ctx iris.Context) *Context {
	return &Context{
		irisContext: ctx,
		isRender:    false,
		session:     sess.Start(ctx),
	}
}

type Context struct {
	irisContext iris.Context

	ControllerName_ string
	ActionName_     string

	isRender bool
	session  *sessions.Session
}

func (ctx *Context) ControllerName() string {
	return ctx.ControllerName_
}

func (ctx *Context) ActionName() string {
	return ctx.ActionName_
}

func (ctx *Context) Params(name string) string {
	return ctx.irisContext.Params().Get(name)
}

func (ctx *Context) SetSession(key, value string) {
	// TODO: add session options, 例如 有效时间等
	ctx.session.Set(key, value)
}

func (ctx *Context) Session(key string) string {
	return ctx.session.GetString(key)
}

func (ctx *Context) RenderText(body string) {
	ctx.irisContext.WriteString(body)
	ctx.isRender = true
}

func (ctx *Context) RenderTextf(format string, values ...interface{}) {
	ctx.irisContext.Writef(format, values...)
	ctx.isRender = true
}

func (ctx *Context) Render(name ...string) {
	if !ctx.isRender {
		ctx.isRender = true
		var templateFilePath string
		// TODO: 获取的view文件的最后可以加上请求的类型 例如, xxx.html 目前都是默认html
		if len(name) <= 0 {
			templateFilePath = filepath.Join(ctx.ControllerName(), ctx.ActionName()+".html")
		} else {
			templateFilePath = filepath.Join(name[0] + ".html")
		}
		logger.Info("渲染html view:" + templateFilePath)
		// TODO: https://github.com/gobuffalo/plush 或者 https://github.com/CloudyKit/jet 模板是比较好的
		// TODO: 开发模式下应该每次都读view文件，生产环境下应该打包到内存中存储起来，这样就不用和html文件导出拷贝了，一个可执行程序就行
	}
}

func (ctx *Context) RedirectTo(url string, statusHeader ...int) {
	ctx.irisContext.Redirect(url, statusHeader...)
}
