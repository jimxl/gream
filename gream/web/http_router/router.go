package http_router

import (
	"github.com/gorilla/mux"
)

var re *mux.Router

func init() {
	re = mux.NewRouter()
}

func Router() *mux.Router {
	return re
}
