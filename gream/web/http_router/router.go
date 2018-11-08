package http_router

import (
	"github.com/gin-gonic/gin"

	"gbs/gream/config"
)

var re *Engine

func init() {
	re = &Engine{gin.Default()}
	// re = &Engine{gin.New()}
	// re.Use(gin.Recovery())
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

func Use(handler HandlerFunc) {}

func GET(path string, handler HandlerFunc) {
	re.GET(path, func(c *gin.Context) {
		handler(&Context{c})
	})
}
