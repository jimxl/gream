package rint

// Gcd 返回最大公约数
func Gcd(x, y uint) uint {
	if x == 0 {
		return y
	}
	if y == 0 {
		return x
	}
	tmp := x % y
	if tmp > 0 {
		return Gcd(y, tmp)
	}
	return y
}
