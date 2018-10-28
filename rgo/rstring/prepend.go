package rstring

import (
	"strings"
)

func prepend(str string, other ...string) string {
	return strings.Join(other, "") + str
}
