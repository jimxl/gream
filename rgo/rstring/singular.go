package rstring

import (
	"github.com/jinzhu/inflection"
)

func Singular(s string) string {
	return inflection.Singular(s)
}
