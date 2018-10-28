package rint

import (
	"math"
)

func Pow(a, b int) int {
	return int(math.Pow(float64(a), float64(b)))
}
