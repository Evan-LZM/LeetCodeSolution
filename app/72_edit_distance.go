package app

func MinDistance(word1, word2 string) int {
	len1 := len(word1)
	len2 := len(word2)
	//len1 is row,len2 is col
	dp := get2DArray(len1+1, len2+1)
	//initial dic
	dp[0][0] = 0
	for i := 0; i <= len1; i++ {
		dp[i][0] = i
	}
	for j := 0; j <= len2; j++ {
		dp[0][j] = j
	}
	for i := 1; i <= len1; i++ {
		for j := 1; j <= len2; j++ {
			if word1[i-1] == word2[j-1] {
				dp[i][j] = dp[i-1][j-1]
			} else {
				dp[i][j] = minDistance(minDistance(dp[i-1][j], dp[i][j-1]), dp[i-1][j-1]) + 1
			}
		}
	}
	return dp[len1][len2]
}

func minDistance(a, b int) int {
	if a > b {
		return b
	}
	return a
}

func get2DArray(m, n int) [][]int {
	var res [][]int
	for i := 0; i < m; i++ {
		res = append(res, make([]int, n))
	}
	return res
}
