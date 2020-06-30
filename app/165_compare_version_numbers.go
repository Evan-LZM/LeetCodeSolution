package app

import (
	"strconv"
	"strings"
)

func CompareVersion(version1 string, version2 string) int {
	v1 := strings.Split(version1, ".")
	v2 := strings.Split(version2, ".")
	leng := min(len(v1), len(v2))
	var res int
	for i := 0; i < leng; i++ {
		num1, _ := strconv.Atoi(v1[i])
		num2, _ := strconv.Atoi(v2[i])
		if num1 == num2 {
			res = 0
			continue
		} else if num1 < num2 {
			return -1
		} else {
			return 1
		}
	}
	if len(v1) > len(v2) {
		value, _ := strconv.Atoi(strings.Join(v1[len(v2):], ""))
		if value == 0 {
			return 0
		}
		return 1
	} else if len(v1) < len(v2) {
		value, _ := strconv.Atoi(strings.Join(v2[len(v1):], ""))
		if value == 0 {
			return 0
		}
		return -1
	}
	return res
}
