package app

func IsCardMatch(s string, p string) bool {
	if len(s) == 0 && len(p) == 0 {
		return true
	}
	if len(s) > 0 && len(p) == 0 {
		return false
	}
	dp := createMatchDip(len(s), len(p))
	dp[0][0] = true
	//initial dp
	for i := 1; i < len(p)+1; i++ {
		if p[i-1] != '*' {
			dp[0][i] = false
		} else {
			dp[0][i] = dp[0][i-1]
		}
	}
	//loop s,p strings
	for i := 1; i <= len(s); i++ {
		for j := 1; j <= len(p); j++ {
			if p[j-1] == '?' || p[j-1] == s[i-1] {
				dp[i][j] = true
			} else if p[j-1] == '*' {
				dp[i][j] = dp[i-1][j] || dp[i][j-1]
			} else {
				dp[i][j] = false
			}
		}
	}
	return dp[len(s)][len(p)]
}

//define dp
func createMatchDip(m int, n int) [][]bool {
	//m is row, n is col
	var dp [][]bool
	for i := 0; i <= m; i++ {
		dp = append(dp, make([]bool, n+1))
	}
	return dp
}
