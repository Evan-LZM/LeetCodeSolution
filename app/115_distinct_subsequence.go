package app

func numDistinct(s string, t string) int {
	if len(s) < len(t) {
		return 0
	}
	dp := create2DDP(len(t)+1, len(s)+1)
	//initial dp
	for i := 0; i <= len(s); i++ {
		dp[0][i] = 1
	}

	for i := 1; i <= len(t); i++ {
		for j := 1; j <= len(s); j++ {
			dp[i][j] += dp[i][j-1]
			if s[j-1] == t[i-1] {
				dp[i][j] += dp[i-1][j-1]
			}
		}
	}
	return dp[len(t)][len(s)]
}

func create2DDP(m, n int) [][]int {
	var data [][]int
	for i := 0; i <= m; i++ {
		data = append(data, make([]int, n+1))
	}
	return data
}
