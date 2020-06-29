package app

func Partition2(s string) [][]string {
	if len(s) == 0 {
		return [][]string{}
	}
	res := [][]string{}
	for i := 1; i <= len(s); i++ {
		if !isPalindrome2(s[:i]) {
			continue
		}
		if i == len(s) {
			res = append(res, []string{s})
		}
		temp := Partition2(s[i:])
		for _, t := range temp {
			n := append([]string{s[:i]}, t...)
			res = append(res, n)
		}
	}
	return res
}

func isPalindrome2(s string) bool {
	if len(s) == 0 {
		return true
	}
	i := 0
	j := len(s) - 1
	for i < j {
		if s[i] != s[j] {
			return false
		}
		i++
		j--
	}
	return true
}
