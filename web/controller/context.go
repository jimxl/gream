package controller

type Context interface {
	ControllerName() string
	ActionName() string

	Params(name string) string
	SetSession(key, value string)
	Session(key string) string
	RenderText(body string)
	RenderTextf(format string, values ...interface{})
	Render(name ...string)
	RedirectTo(url string, statusHeader ...int)

	ToView(name string, value interface{})
}
