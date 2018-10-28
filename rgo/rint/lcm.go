package rint

// Lcm 返回两个正整数的最小公倍数
func Lcm(x, y uint) uint {
	return x * y / Gcd(x, y)
}
