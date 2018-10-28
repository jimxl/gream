package rstring

import (
	"strings"
)

func IsInclude(str string, substr string) bool {
	return strings.Contains(str, substr)
}
