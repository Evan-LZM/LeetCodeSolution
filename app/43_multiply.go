package app

import (
	"bytes"
	"strconv"
)

func Multiply(num1 string, num2 string) string {
	if len(num1) == 0 || len(num2) == 0 || num1 == "0" || num2 == "0" {
		return "0"
	}
	result := make([]int, len(num1)+len(num2))
	for i, value := range num1 {
		for j, k := range num2 {
			result[len(num1)+len(num2)-i-j-2] += int(value-'0') * int(k-'0')
		}
	}
	carry := 0
	for index, value := range result {
		t := carry + value
		result[index] = t % 10
		carry = t / 10
	}
	var buf bytes.Buffer
	for i := len(result) - 1; i >= 0; i-- {
		if !(result[i] == 0 && len(result)-1 == i) {
			buf.WriteString(strconv.Itoa(result[i]))
		}
	}
	ans := buf.String()
	if len(ans) == 0 {
		return "0"
	}
	return ans
}
