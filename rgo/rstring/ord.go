package rstring

func Ord(str string) int {
	if len(str) == 0 {
		return -1
	}
	return int(str[0])
}
