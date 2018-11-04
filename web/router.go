package web

import (
	. "gbs/gream/web/router"
)

func Draw() {

	GET("/home_json/{name}", "home#index_json")
	GET("/test/{name}", "scope/home#index")

	// scope := r.Scope("/scope")
	// {
	// 	scope.GET("/home_json/:name", "home#index")
	// }

	// namespace := r.Namespace("/admin")
	// {
	// 	namespace.GET("/home/:name", "home#index_text")
	// }

	// r.Resources("users")
}
