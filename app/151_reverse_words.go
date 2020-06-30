package app

import (
	"strings"
)

func ReverseWords(s string) string {
	list := strings.Fields(s)
	var res []string
	for i := len(list) - 1; i >= 0; i-- {
		res = append(res, list[i])
	}
	return strings.Join(res, " ")
}
