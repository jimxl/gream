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

	methods := map[string]bool{
		"index":   true,
		"new":     true,
		"create":  true,
		"show":    true,
		"edit":    true,
		"update":  true,
		"destroy": true,
	}

	if len(opts) >= 1 {
		if only, ok := opts[0]["only"]; ok {
			delete(opts[0], "only")
			methods = map[string]bool{
				"index":   false,
				"new":     false,
				"create":  false,
				"show":    false,
				"edit":    false,
				"update":  false,
				"destroy": false,
			}
			for _, ms := range rstring.Split(only, ",") {
				methods[ms] = true
			}
		} else if except, ok := opts[0]["except"]; ok {
			delete(opts[0], "except")
			methods = map[string]bool{
				"index":   true,
				"new":     true,
				"create":  true,
				"show":    true,
				"edit":    true,
				"update":  true,
				"destroy": true,
			}
			for _, ms := range rstring.Split(except, ",") {
				methods[ms] = false
			}
		}
	}

	for key, ok := range methods {
		if ok {
			resourcesMethods[key](pluralName, r, opts)
		}
	}
}
