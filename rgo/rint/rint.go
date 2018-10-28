package rint

// IsNegative 判断是否为负数，小于零的时候返回true
func IsNegative(i int) bool {
	return i < 0
}

// IsPositive 判断是否为正数, 大于0的时候返回true
func IsPositive(i int) bool {
	return i > 0
}

// IsNonzero 判断数字是否为0
func IsNonzero(i int) bool {
	return i == 0
}
