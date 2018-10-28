package rstring

import (
	"strings"
)

func EachLine(str string, f func(int, string)) {
	lines := strings.Split(str, "\n")

	for index, line := range lines {
		f(index, line)
	}
}
