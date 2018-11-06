package web

import (
	. "gbs/gream/web/router"
)

func Draw() {

	GET("/home_json/{name}", "home#index_json")
	GET("/test/{name}", "scope/home#index")

	scope := Scope("scope")
	{
		scope.GET("/home_json/{name}", "home#index")
	}

	namespace := Namespace("admin")
	{
		namespace.GET("/home/{name}", "home#index")
	}

	// r.Resources("users")
}
