package rstring

import (
	"regexp"
)

// TODO: 类似ruby可以添加末尾一个block
func Gsub(str string, pattern string, replacement string) string {
	re := regexp.MustCompile(pattern)
	return re.ReplaceAllString(str, replacement)
}
