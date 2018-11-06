package middleware

import (
	"fmt"
	"net/http"
	"time"

	. "gbs/gream/web/response"
)

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
