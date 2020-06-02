package app

import "strconv"

func countAndSay(n int) string {
	if n == 1 {
		return "1"
	}
	cur := "1"
	for c := 2; c <= n; c++ {
		cur = getNext(cur)
	}
	return cur
}

func getNext(s string) string {
	i := 0
	res := ""
	for i < len(s) {
		j := i + 1
		for j < len(s) {
			if s[j] != s[j-1] {
				break
			} else {
				j++
			}
		}
		res = res + strconv.Itoa(j-i) + string(s[i])
		i = j
	}
	return res
}
