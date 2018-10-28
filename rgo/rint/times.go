package rint

func Times(count int, f func(index int)) {
	for i := 0; i < count; i++ {
		f(i)
	}
}
