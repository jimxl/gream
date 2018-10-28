package rstring

import (
	"strings"
)

// TODO: 支持正则
func Index(str string, substr string) int {
	return strings.Index(str, substr)
}
