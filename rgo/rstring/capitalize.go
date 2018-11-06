package rstring

import (
	"github.com/huandu/xstrings"
)

func Capitalize(str string) string {
	return xstrings.ToCamelCase(str)
}
