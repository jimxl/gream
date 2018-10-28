package rint

func Gcdlcm(x, y uint) (gcd uint, lcm uint) {
	gcd = Gcd(x, y)
	lcm = x * y / gcd
	return
}
