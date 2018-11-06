package router

import (
	"fmt"
	"net/http"
	"strings"

	"gbs/gream/config"
	"gbs/gream/web/http_router"

	"github.com/gorilla/mux"
)

var re *mux.Router

func init() {
	re = http_router.Router()
}

func Run() {

	srv := &http.Server{
		Handler:      re,
		Addr:         config.App.Addr,
		WriteTimeout: config.App.WriteTimeout,
		ReadTimeout:  config.App.ReadTimeout,
	}

	printUrls()
	srv.ListenAndServe()
}

func GET(path string, controllerAndAction string) *GScope {
	scope := &GScope{
		route: re.NewRoute(),
	}
	return scope.GET(path, controllerAndAction)
}

func Scope(path string) *GScope {
	scope := &GScope{
		route: re.NewRoute(),
	}
	return scope.Scope(path)
}

func Namespace(path string) *GScope {
	scope := &GScope{
		route: re.NewRoute(),
	}
	return scope.Namespace(path)
}

func printUrls() {
	err := re.Walk(func(route *mux.Route, router *mux.Router, ancestors []*mux.Route) error {
		pathTemplate, err := route.GetPathTemplate()
		if err == nil {
			fmt.Println("ROUTE:", pathTemplate)
		}
		pathRegexp, err := route.GetPathRegexp()
		if err == nil {
			fmt.Println("Path regexp:", pathRegexp)
		}
		queriesTemplates, err := route.GetQueriesTemplates()
		if err == nil {
			fmt.Println("Queries templates:", strings.Join(queriesTemplates, ","))
		}
		queriesRegexps, err := route.GetQueriesRegexp()
		if err == nil {
			fmt.Println("Queries regexps:", strings.Join(queriesRegexps, ","))
		}
		methods, err := route.GetMethods()
		if err == nil {
			fmt.Println("Methods:", strings.Join(methods, ","))
		}
		fmt.Println()
		return nil
	})
	if err != nil {
		fmt.Println(err)
	}
}
