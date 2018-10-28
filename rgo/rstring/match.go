package rstring

import (
	"regexp"
)

func Match(str, pattern string) (bool, []string) {
	re := regexp.MustCompile(pattern)
	strs := re.FindAllString(str, -1)
	return len(strs) != 0, strs
}
