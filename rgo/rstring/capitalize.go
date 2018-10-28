package rstring

import (
	"strings"
)

func Capitalize(str string) string {
	return strings.Title(Downcase(str))
}
