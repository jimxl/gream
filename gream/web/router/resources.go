package router

import "gbs/rgo/rstring"

var resourcesMethods = map[string]func(string, *Router, []H){
	"index": func(pluralName string, r *Router, opts []H) {
		r.GET(urlJoin("/", pluralName), pluralName+"#index", opts...)
	},
	"new": func(pluralName string, r *Router, opts []H) {
		r.GET(urlJoin("/", pluralName, "new"), pluralName+"#new", opts...)
	},
	"create": func(pluralName string, r *Router, opts []H) {
		r.POST(urlJoin("/", pluralName), pluralName+"#create", opts...)
	},
	"show": func(pluralName string, r *Router, opts []H) {
		r.GET(urlJoin("/", pluralName, "{id}"), pluralName+"#show", opts...)
	},
	"edit": func(pluralName string, r *Router, opts []H) {
		r.GET(urlJoin("/", pluralName, "{id}/edit"), pluralName+"#edit", opts...)
	},
	"update": func(pluralName string, r *Router, opts []H) {
		r.PUT(urlJoin("/", pluralName, "{id}"), pluralName+"#update", opts...)
		r.PATCH(urlJoin("/", pluralName, "{id}"), pluralName+"#update", opts...)
	},
	"destroy": func(pluralName string, r *Router, opts []H) {
		r.DELETE(urlJoin("/", pluralName, "{id}"), pluralName+"#destroy", opts...)
	},
}

func (r *Router) Resources(name string, opts ...H) {
	pluralName := rstring.Plural(name)

	for _, method := range resourcesMethods {
		method(pluralName, r, opts)
	}
}
