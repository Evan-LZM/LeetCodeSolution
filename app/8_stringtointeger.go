package app

import (
	"math"
	"strconv"
)

func MyAtoi(str string) int {
	var i, num int //i store the current index of the str
	// num store the number total value
	sign := 1
	if str == "" {
		return 0
	}
	for i < len(str) && str[i] == 32 {
		i++
	}
	if i >= len(str) {
		return 0
	}
	if str[i] == 43 {
		i++
	} else if str[i] == 45 {
		sign = -1
		i++
	}
	for ; i < len(str); i++ {
		if str[i] != 0 && (str[i] < 48 || str[i] > 57) {
			return num * sign
		}
		n, _ := strconv.Atoi(string(str[i]))
		num = num*10 + n
		if num*sign > math.MaxInt32 {
			return math.MaxInt32
		} else if num*sign < math.MinInt32 {
			return math.MinInt32
		}
	}
	return num * sign
}
