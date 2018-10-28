package rstring

import (
	"strings"
	"unicode"
)

func Lstrip(str string) string {
	return strings.TrimLeftFunc(str, unicode.IsSpace)
}
