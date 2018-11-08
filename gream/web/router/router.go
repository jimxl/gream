package router

var scope = &GScope{}

func GET(path string, controllerAndAction string) {
	scope.GET(path, controllerAndAction)
}

// func Scope(path string) *GScope {
// 	return scope.Scope(path)
// }

// func Namespace(path string) *GScope {
// 	return scope.Namespace(path)
// }
