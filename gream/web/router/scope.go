package router

// // Resources TODO: 要实现Only和Expect方法，用于关闭某些默认路由
// // TODO:  由于gin不支持 GET /users/new  和 GET /users/:id 这样的路由，会冲突只能换成 GET /user/:id
// func (scope *Scope) Resources(name string, opts ...Opt) *Scope {
// 	var actions map[string]bool
// 	if len(opts) == 1 {
// 		opt := opts[0]
// 		actions = rarray.OnlyAndExpect(opt.Only, opt.Expect, resourcesActions)
// 	}

// 	pluralName := rstring.Plural(name)
// 	singularName := rstring.Singular(name)

// 	if len(actions) == 0 || actions["index"] {
// 		scope.GET("/"+pluralName, pluralName+"#index")
// 	}
// 	if len(actions) == 0 || actions["new"] {
// 		scope.GET("/"+singularName+"/new", pluralName+"#new")
// 	}
// 	if len(actions) == 0 || actions["create"] {
// 		scope.POST("/"+name, pluralName+"#create")
// 	}
// 	if len(actions) == 0 || actions["show"] {
// 		scope.GET("/"+pluralName+"/:id", pluralName+"#show")
// 	}
// 	if len(actions) == 0 || actions["edit"] {
// 		scope.GET("/"+singularName+"/edit", pluralName+"#edit")
// 	}
// 	if len(actions) == 0 || actions["update"] {
// 		scope.PATCH("/"+pluralName+"/:id", pluralName+"#update")
// 		scope.PUT("/"+pluralName+"/:id", pluralName+"#update")
// 	}
// 	if len(actions) == 0 || actions["destroy"] {
// 		scope.DELETE("/"+pluralName+"/:id", pluralName+"#destroy")
// 	}
// 	return &Scope{
// 		basePath: urlPath.Join(scope.basePath, name),
// 	}
// }
