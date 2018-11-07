package web

import (
	. "gbs/gream/web/router"
)

func Draw() {

	GET("/home/{name}", "home#index")
	GET("/home_json/{name}", "home#index_json")
	GET("/scope_home/{name}", "scope/home#index")

	scope := Scope("scope")
	{
		scope.GET("/home1/{name}", "home#index")
		scope.GET("/home2/{name}", "home#index")
	}

	namespace := Namespace("admin")
	{
		namespace.GET("/homea/{name}", "home#index")
		namespace.GET("/homeb/{name}", "home#index")
	}

	// r.Resources("users")
}
