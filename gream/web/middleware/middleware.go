package middleware

import (
	"fmt"
	"gbs/gream/logger"
	"gbs/gream/web/http_router"
	. "gbs/gream/web/response"
	"time"

	"net/http"

	"github.com/gorilla/mux"
)

var re *mux.Router

func init() {
	re = http_router.Router()
	re.Use(responseMiddleware)
	Use(loggerMiddleWare)
}

func responseMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		response := &Response{ResponseWriter: w}
		next.ServeHTTP(response, r)
	})
}

type HandlerFunc func(response *Response, request *http.Request)

func (f HandlerFunc) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	f(w.(*Response), r)
}

func Use(mwf ...mux.MiddlewareFunc) {
	re.Use(mwf...)
}

func loggerMiddleWare(next http.Handler) http.Handler {
	return HandlerFunc(func(response *Response, r *http.Request) {
		t := time.Now()

		requestMethod := r.Method
		url := r.RequestURI

		logger.Info(
			fmt.Sprintf(
				"Started %s \"%s\" for %s at %s",
				requestMethod,
				url,
				r.RemoteAddr,
				t.Format("2006-01-02 15:04:05 +0800"),
			))

		next.ServeHTTP(response, r)

		latency := time.Since(t)
		status := response.StatusCode()

		logger.Info(
			fmt.Sprintf(
				"Completed %s %v in %v",
				http.StatusText(status),
				status,
				latency,
			))

	})
}
