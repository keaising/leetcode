package pow

import "math"

func myPow(x float64, n int) float64 {
	return math.Pow(float64(x), float64(n))
}
