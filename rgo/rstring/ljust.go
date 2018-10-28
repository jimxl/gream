package rstring

import (
	"strings"
)

func Ljust(str string, length int, padstr ...string) string {
	strLen := len(str)
	if length <= strLen {
		return str[:length]
	}

	pad := " "
	if len(padstr) >= 1 {
		pad = padstr[0]
	}

	retStr := strings.Repeat(pad, length-strLen)
	return str + retStr
}
