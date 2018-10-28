package router

import (
	"gbs/gream/config"

	"github.com/gin-gonic/gin"
)

var re = &engine{router: gin.New()}

type engine struct {
	router *gin.Engine
}

func (e *engine) run() {
	e.router.Run(config.App.Host + ":" + config.App.Port)
}
