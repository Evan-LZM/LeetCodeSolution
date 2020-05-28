package app

//IsMatch '.' Matches any single character.
//IsMatch '*' Matches zero or more of the preceding element.
func IsMatch(s string, p string) bool {
	if len(s) == 0 && len(p) == 0 {
		return true
	}
	if len(s) > 0 && len(p) == 0 {
		return false
	}
	dp := createRegularDP(len(s), len(p))
	dp[0][0] = true
	//initial dp
	for i := 1; i <= len(p); i++ {
		if p[i-1] == '*' {
			dp[0][i] = dp[0][i-2]
		}
	}
	for i := 1; i <= len(s); i++ {
		for j := 1; j <= len(p); j++ {
			if p[j-1] == '.' || s[i-1] == p[j-1] {
				dp[i][j] = dp[i-1][j-1]
			} else if p[j-1] == '*' {
				dp[i][j] = dp[i][j-2] || dp[i][j]
				if s[i-1] == p[j-2] || p[j-2] == '.' {
					//保留原来的真。避免被i-1和j-1冲掉
					dp[i][j] = dp[i][j] || dp[i-1][j]

					dp[i][j] = dp[i][j] || dp[i][j-1]

				}
			}
		}
	}
	return dp[len(s)][len(p)]
}

func createRegularDP(m, n int) [][]bool {
	var dp [][]bool
	for i := 0; i <= m; i++ {
		dp = append(dp, make([]bool, n+1))
	}
	return dp
}
