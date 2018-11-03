package router

import (
	"gbs/gream/config"
	"net/http"

	"github.com/gorilla/mux"
)

var re = &engine{router: mux.NewRouter()}

type engine struct {
	router *mux.Router
}

func (e *engine) run() {

	srv := &http.Server{
		Handler:      e.router,
		Addr:         config.App.Addr,
		WriteTimeout: config.App.WriteTimeout,
		ReadTimeout:  config.App.ReadTimeout,
	}

	srv.ListenAndServe()
}

func Get(path string, controllerAndAction string) {
	re.router.Methods("GET").HandlerFunc(path)
}
