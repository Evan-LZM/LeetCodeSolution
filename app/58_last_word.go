package app

import (
	"strings"
)

func LengthOfLastWord(s string) int {
	s = strings.TrimRight(s, " ")
	array := strings.Split(s, " ")
	return len(array[len(array)-1])
}

/*
other solution:
 	if len(s) == 0 {
		return 0
	}

	last := len(s) - 1
	for ; last >= 0; last-- {
		if s[last] != ' ' {
			break
		}
	}

	if last < 0 {
		return 0
	}

	first := last
	for ; first >= 0; first-- {
		if s[first] == ' ' {
			break
		}
	}

	return last - first
*/
