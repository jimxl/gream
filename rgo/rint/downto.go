package rint

// DownTo 递减的迭代
func DownTo(from int, to int, block func(i int)) {
	if to > from {
		return
	}
	for i := from; i >= to; i-- {
		block(i)
	}
}
