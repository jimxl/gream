package Response

import "net/http"

type Response struct {
	http.ResponseWriter
	statusCode int
}

func (r *Response) WriteHeader(status int) {
	r.statusCode = status
}

func (r *Response) StatusCode() int {
	return r.statusCode
}
