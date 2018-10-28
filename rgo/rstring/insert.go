package rstring

func Insert(str string, index int, otherstr string) string {
	return str[:index] + otherstr + str[index:]
}
