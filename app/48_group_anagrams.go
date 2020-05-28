package app

import (
	"sort"
	"strings"
)

func GroupAnagrams(strs []string) [][]string {
	var result [][]string
	if len(strs) == 0 {
		return result
	}
	dic := make(map[string][]string)
	for _, value := range strs {
		s := sortString(value)
		dic[s] = append(dic[s], string(value))
	}
	for _, value := range dic {
		result = append(result, value)
	}
	return result
}

func sortString(s string) string {
	temp := strings.Split(s, "")
	sort.Strings(temp)
	return strings.Join(temp, "")
}
