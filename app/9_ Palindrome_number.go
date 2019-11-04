package app

import (
	"bytes"
	"fmt"
	"strconv"
	"strings"
)

func IsPalindrome(x int) bool {
	if x < 0 {
		return false
	}
	s := strconv.Itoa(x)
	data := ""
	for i := len(s) - 1; i >= 0; i-- {
		data += string(s[i])
	}
	fmt.Println(s)
	fmt.Println(data)
	if strings.Compare(s, data) == 0 {
		return true
	} else {
		return false
	}
}

func IsPalindromeV2(x int) bool {
	if x < 0 {
		return false
	}
	var buf bytes.Buffer
	temp := x
	for temp > 0 {
		fmt.Println(temp)
		buf.WriteRune(rune(temp%10 + '0'))
		temp /= 10
	}
	fmt.Println(buf.String())
	v, _ := strconv.Atoi(buf.String())
	return v == x
}

func IsPalindromeV3(x int) bool {
	if x < 0 {
		return false
	}
	v := x
	var result string
	for v > 0 {
		result += strconv.Itoa(v % 10)
		v /= 10
	}
	a, _ := strconv.Atoi(result)
	return x == a
}
