package http_router

import (
	"github.com/gin-gonic/gin"

	"gbs/gream/config"
)

var re *Engine

func init() {
	//re = &Engine{gin.Default()}
	re = &Engine{gin.New()}
	re.Use(gin.Recovery())
}

func Router() *Engine {
	return re
}

func Run() {
	re.Run(config.App.Addr)
}

type Engine struct {
	*gin.Engine
}

type Context struct {
	*gin.Context
}

type H = gin.H
type HandlerFunc = func(c *Context)

func Use(_ HandlerFunc) {
//	#TODO: 加上中间件的使用
}

func GET(path string, handler HandlerFunc) {
	re.GET(path, func(c *gin.Context) {
		handler(&Context{c})
	})
}

func POST(path string, handler HandlerFunc) {
	re.POST(path, func(c *gin.Context) {
		handler(&Context{c})
	})
}

func DELETE(path string, handler HandlerFunc) {
	re.DELETE(path, func(c *gin.Context) {
		handler(&Context{c})
	})
}

func PATCH(path string, handler HandlerFunc) {
	re.PATCH(path, func(c *gin.Context) {
		handler(&Context{c})
	})
}

func PUT(path string, handler HandlerFunc) {
	re.PUT(path, func(c *gin.Context) {
		handler(&Context{c})
	})
}

func OPTIONS(path string, handler HandlerFunc) {
	re.OPTIONS(path, func(c *gin.Context) {
		handler(&Context{c})
	})
}

func HEAD(path string, handler HandlerFunc) {
	re.HEAD(path, func(c *gin.Context) {
		handler(&Context{c})
	})
}

func Any(path string, handler HandlerFunc) {
	re.Any(path, func(c *gin.Context) {
		handler(&Context{c})
	})
}
