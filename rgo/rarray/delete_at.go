package rarray

func DeleteAt(array []interface{}, index int) []interface{} {
	copy(array[index:], array[index+1:])
	return array[:len(array)-1]
}
