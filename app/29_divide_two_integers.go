package app

import (
	"math"
)

//Divide a
func Divide(dividend int, divisor int) int {
	if dividend == 0 {
		return 0
	}
	if dividend == (math.MinInt32) && divisor == -1 {
		return math.MaxInt32
	}
	u := dividend > 0
	v := divisor > 0
	positive := (u && v) || (!u && !v)
	result := 0
	if !u {
		dividend = -dividend
	}
	if !v {
		divisor = -divisor
	}
	for dividend > 0 {
		q := 0
		dividend, q = helperdivide(dividend, divisor)
		result += q
	}

	if positive {
		return result
	}
	return result
}

func helperdivide(dividend, divisor int) (int, int) {
	if dividend < divisor {
		return 0, 0
	}
	q := 1
	for (divisor << 1) < dividend {
		q = q << 1
		divisor = divisor << 1
	}
	return dividend - divisor, q
}
