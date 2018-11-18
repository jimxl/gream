package web

import (
	. "gbs/gream/web/router"
)

func init() {
	ApplicationRouterDraw = func() {

		GET("/home/{name}", "home#index")
		GET("/home_json/{name}", "home#index_json")
		GET("/admin_home/{name}", "admin/home#index")

		scope := Scope("scope")
		{
			scope.GET("/home1/{name}", "home#index")
			scope.GET("/home2/{name}", "home#index")
		}

		scope = Scope(H{"module": "admin"})
		{
			scope.GET("/home1/{name}", "home#index")
			scope.GET("/home2/{name}", "home#index")
		}

		GET("/home_path/{name}", "home#index", H{"path": "ttt"})
		GET("/home_module/{name}", "home#index", H{"module": "admin"})

		namespace := Namespace("admin")
		{
			namespace.GET("/homea/{name}", "home#index")
			namespace.GET("/homeb/{name}", "home#index")
		}

		Resources("users")
	}

}
