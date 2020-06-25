package app

import "strings"

func isPalindrome(s string) bool {
	s = strings.ToUpper(s)
	i := 0
	j := len(s) - 1
	for i < j {
		for i < j && !((s[i] >= '0' && s[i] <= '9') || (s[i] >= 'a' && s[i] <= 'z') || (s[i] >= 'A' && s[i] <= 'Z')) {
			i++
		}
		for i < j && !((s[j] >= '0' && s[j] <= '9') || (s[j] >= 'a' && s[j] <= 'z') || (s[j] >= 'A' && s[j] <= 'Z')) {
			j--
		}
		if i < j && s[i] != s[j] {
			return false
		}
		i++
		j--
	}
	return true
}
