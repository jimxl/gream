# Gream

Gream is a Dream Web Framework written in Go(Golang)

## Router

```go
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

//Resources("users", H{"except": "index"})
Resources("users", H{"only": "index,new"})
```

## Controller and Action

```go
package controllers

import (
	"fmt"
	. "github.com/jimxl/gream/web"
)

func init() {
	Controller("home")

	BeforeAll(allFilter)
	BeforeAction("index", indexFilter)

	Action("index", func(ctx Context) {
		ctx.Writef("hello, %s", ctx.Params().GetString("name"))
	})
}

func indexFilter(ctx Context) bool {
	fmt.Printf("index filter, controller %s, action: %s\n", ctx.Values().Get("controller"), ctx.Values().Get("action"))
	return true
}

func allFilter(ctx Context) bool {
	fmt.Printf("all filter, controller %s, action: %s\n", ctx.Values().Get("controller"), ctx.Values().Get("action"))
	return true
}
```
