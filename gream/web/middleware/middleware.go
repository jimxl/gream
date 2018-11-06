package middleware

import (
	"gbs/gream/web/http_router"
	. "gbs/gream/web/response"

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
