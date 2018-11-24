package middleware

//
//import (
//	"fmt"
//	"gbs/gream/web/http_router"
//	"net/http"
//	"time"
//
//	"gbs/gream/logger"
//)
//
//func loggerMiddleWare() http_router.HandlerFunc {
//	return func(c *http_router.Context) {
//		t := time.Now()
//
//		requestMethod := c.Request.Method
//		url := c.Request.RequestURI
//
//		logger.Info(
//			fmt.Sprintf(
//				"Started %s \"%s\" for %s at %s",
//				requestMethod,
//				url,
//				c.ClientIP(),
//				t.Format("2006-01-02 15:04:05 +0800"),
//			))
//
//		c.Next()
//
//		latency := time.Since(t)
//		status := c.Writer.Status()
//
//		logger.Info(
//			fmt.Sprintf(
//				"Completed %s %v in %v\n\n",
//				http.StatusText(status),
//				status,
//				latency,
//			))
//
//	}
//}
