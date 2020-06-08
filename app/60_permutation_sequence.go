package app

import (
	"strconv"
	"strings"
)

//3,3 213
func GetPermuation(n int, k int) string {
	var result strings.Builder
	ints := make([]int, n)
	for i := 0; i < n; i++ {
		ints[i] = i + 1
	}
	for len(ints) > 0 {
		f := permuRecur(len(ints) - 1)
		m := (k - 1) / f
		k = k - f*m
		result.WriteString(strconv.Itoa(ints[m]))
		ints = append(ints[:m], ints[m+1:]...)
	}
	return result.String()
}

func permuRecur(n int) int {
	var sum = 1
	for i := 1; i <= n; i++ {
		sum *= i
	}
	return sum
}
