package http_router

func (ctx *Context) SetSession(key, value string) {
	// TODO: add session options, 例如 有效时间等
	ctx.irisContext.SetCookieKV(key, value)
}

func (ctx *Context) Session(key string) string {
	return ctx.irisContext.GetCookie(key)
}
