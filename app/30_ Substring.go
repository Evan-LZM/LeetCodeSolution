package app

import "fmt"

//FindSubstring 30
func FindSubstring(s string, words []string) []int {
	var result []int
	if len(s) == 0 {
		return result
	}
	if len(words) == 0 || len(words[0]) == 0 {
		return result
	}

	dict := make(map[string]int)
	for _, word := range words {
		if o, ok := dict[word]; ok {
			dict[word] = o + 1
		} else {
			dict[word] = 1
		}
	}
	l := len(words[0])
	for i := 0; i <= len(s)-l*len(words); i++ {
		if compareSubString(s[i:i+len(words)*l], dict, l) {
			result = append(result, i)
		}
	}
	return result
}

func compareSubString(s string, dict map[string]int, l int) bool {
	d := make(map[string]int)
	fmt.Println(s)
	for i := 0; i < len(s); i += l {
		sub := s[i : i+l]
		occurance := get0occurrence(dict, sub)
		m := get0occurrence(d, sub)
		d[sub] = m + 1
		if m+1 > occurance {
			return false
		}
	}
	return true
}

func get0occurrence(dict map[string]int, s string) int {
	if occurence, ok := dict[s]; ok {
		return occurence
	}
	return 0
}
