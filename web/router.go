package web

import (
	router "github.com/jimxl/gream/web/router"
)

func Draw() {
	r := router.Create()

	r.GET("/home_json/:name", "home#index_json")
	r.GET("/test/:name", "scope/home#index")

	scope := r.Scope("/scope")
	{
		scope.GET("/home_json/:name", "home#index")
	}

	namespace := r.Namespace("/admin")
	{
		namespace.GET("/home/:name", "home#index_text")
	}

	r.Resources("users")
}
