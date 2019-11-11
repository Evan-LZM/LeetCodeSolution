package app

func strStr(haystack string, needle string) int {
	//method-1
	// if needle == "" {
	// 	return 0
	// }
	// return strings.Index(haystack, needle)

	//method-2
	for i := 0; i <= len(haystack)-len(needle); i++ {
		if compare(haystack[i:i+len(needle)], needle) {
			return i
		}
	}
	return -1
}

func compare(a string, b string) bool {
	if len(a) != len(b) {
		return false
	}
	for i, _ := range a {
		if a[i] != b[i] {
			return false
		}
	}
	return true
}
