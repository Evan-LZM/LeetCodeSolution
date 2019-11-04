package app

func RomanToInt(s string) int {
	dict := map[byte]int{
		'I': 1,
		'V': 5,
		'X': 10,
		'L': 50,
		'C': 100,
		'D': 500,
		'M': 1000,
	}
	result := 0
	for k, _ := range s {
		if k != (len(s)-1) && (dict[s[k]] < dict[s[k+1]]) {
			result -= dict[s[k]]
		} else {
			result += dict[s[k]]
		}
	}
	return result
}
