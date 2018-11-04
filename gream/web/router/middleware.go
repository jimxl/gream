package router

import (
	"fmt"
	"net/http"
	"time"

	"gbs/gream/logger"
)

func init() {
	re.Use(loggerMiddleWare)
}

func loggerMiddleWare(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
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

		next.ServeHTTP(w, r)

		latency := time.Since(t)
		status := w.Header().Get("Status Code")

		logger.Info(
			fmt.Sprintf(
				"Completed %s in %v",
				// http.StatusText(status),
				status,
				latency,
			))

	})
}
