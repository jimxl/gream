package router

import (
	"fmt"
	"net/http"
	"time"

	"gbs/gream/logger"

	"github.com/gin-gonic/gin"
)

func init() {
	gin.SetMode(gin.ReleaseMode)
	re.router.Use(loggerMiddleWare())
	re.router.Use(gin.Recovery())
}

func loggerMiddleWare() gin.HandlerFunc {
	return func(c *gin.Context) {
		t := time.Now()

		requestMethod := c.Request.Method
		url := c.Request.URL.RequestURI()

		logger.Info(
			fmt.Sprintf(
				"Started %s \"%s\" for %s at %s",
				requestMethod,
				url,
				c.ClientIP(),
				t.Format("2006-01-02 15:04:05 +0800"),
			))

		c.Next()
		latency := time.Since(t)
		status := c.Writer.Status()

		logger.Info(
			fmt.Sprintf(
				"Completed %d %s in %v",
				status,
				http.StatusText(status),
				latency,
			))

	}
}
