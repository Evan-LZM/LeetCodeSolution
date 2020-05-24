package app

//GenerateParenthesis aa
func GenerateParenthesis(n int) []string {
	var result []string
	//think of ( as +1 )as -1
	generatehelper("", 0, 0, n, &result)
	return result
}

//u is number of ( , v is number of )
func generatehelper(s string, u int, v int, n int, result *[]string) {
	if u == n && v == n {
		*result = append(*result, s)
		return
	}
	if u == n {
		generatehelper(s+")", u, v+1, n, result)
		return
	}
	if u == v {
		generatehelper(s+"(", u+1, v, n, result)
		return
	}
	generatehelper(s+")", u, v+1, n, result)
	generatehelper(s+"(", u+1, v, n, result)
}
