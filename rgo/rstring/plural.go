package rstring

import (
	"github.com/jinzhu/inflection"
)

func Plural(s string) string {
	return inflection.Plural(s)
}
