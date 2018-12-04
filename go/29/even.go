package enen

import (
	"math"
)

func divide(dividend int, divisor int) int {
	ret := float64(dividend) / float64(divisor)
	t := 0
	if ret > 0 {
		t = int(math.Floor(ret))
	} else {
		t = int(math.Ceil(ret))
	}
	if t > 2147483648 || t < -2147483647 {
		return 2147483647
	}
	return t
}
