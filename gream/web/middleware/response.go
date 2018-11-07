package middleware

import (
	. "gbs/gream/web/response"
	"net/http"
)

func responseMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		response := &Response{ResponseWriter: w}
		next.ServeHTTP(response, r)
	})
}
