package atoi

import (
	"math"
)

func myAtoi(str string) int {
	ret, sign, length, i := 0, 1, len(str), 0
	for i < length && (str[i] == ' ' || str[i] == '\t') {
		i++
	}
	if i < length {
		if str[i] == '+' {
			sign = 1
			i++
		} else if str[i] == '-' {
			sign = -1
			i++
		}
	}
	for i < length && str[i] >= '0' && str[i] <= '9' {
		ret = ret*10 + int(str[i]) - '0'
		if sign*ret > math.MaxInt32 {
			return math.MaxInt32
		} else if sign*ret < math.MinInt32 {
			return math.MinInt32
		}
		i++
	}
	return ret * sign
}
