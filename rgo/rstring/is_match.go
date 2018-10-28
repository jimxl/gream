package rstring

import "regexp"

func IsMatch(str, pattern string) bool {
	re := regexp.MustCompile(pattern)

	return re.MatchString(str)
}
