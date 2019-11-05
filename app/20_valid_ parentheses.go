package app

import (
	"strings"
)

func IsValid(s string) bool {
	if s == "" {
		return true
	}
	if len(s)%2 != 0 {
		return false
	}
	if !strings.ContainsRune(s, 41) && !strings.ContainsRune(s, 93) && !strings.ContainsRune(s, 125) {
		return false
	}
	stack := make([]rune, len(s))
	j := 0

	for _, v := range s {
		if v == 40 || v == 91 || v == 123 {
			stack[j] = v
			j++
		} else if v == 41 || v == 93 || v == 125 {
			if j-1 < 0 {
				return false
			} else if stack[j-1]+1 == v || stack[j-1]+2 == v {
				j--
				continue
			} else {
				return false
			}
		}
	}
	if j != 0 {
		return false
	}
	return true
}
