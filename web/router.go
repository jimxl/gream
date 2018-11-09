package web

import (
	. "gbs/gream/web/router"
)

func Draw() {

	GET("/home/:name", H{"to": "home#index"})
	GET("/home_json/:name", H{"to": "home#index_json"})
	GET("/scope_home/:name", H{"to": "admin/home#index"})

	scope := Scope("scope")
	{
		scope.GET("/home1/:name", H{"to": "home#index"})
		scope.GET("/home2/:name", H{"to": "home#index"})
	}

	namespace := Namespace("admin")
	{
		namespace.GET("/homea/:name", H{"to": "home#index"})
		namespace.GET("/homeb/:name", H{"to": "home#index"})
	}

	// r.Resources("users")
}
