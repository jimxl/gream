package middleware

import "net/http"
import . "gbs/gream/web/response"

func responseMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		response := &Response{ResponseWriter: w}
		next.ServeHTTP(response, r)
	})
}
