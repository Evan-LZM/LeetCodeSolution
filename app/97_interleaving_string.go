package app

/*IsInterleave Given s1, s2, s3, find whether
s3 is formed by the interleaving of s1 and s2.
*/
func IsInterleave(s1, s2, s3 string) bool {
	if len(s1)+len(s2) != len(s3) {
		return false
	}
	dp := createTwoDimentionalArray(len(s1)+1, len(s2)+1)
	dp[0][0] = true
	//initialise dp
	for i := 1; i < len(s1)+1; i++ {
		dp[i][0] = dp[i-1][0] && s1[i-1] == s3[i-1]
	}
	for j := 1; j < len(s2)+1; j++ {
		dp[0][j] = dp[0][j-1] && s2[j-1] == s3[j-1]
	}
	for i := 1; i <= len(s1); i++ {
		for j := 1; j <= len(s2); j++ {
			dp[i][j] = (dp[i-1][j] && s1[i-1] == s3[i+j-1]) || (dp[i][j-1] && s2[j-1] == s3[j-1+i])
		}
	}
	return dp[len(s1)][len(s2)]
}

//createTwoDimentionalArray m is row, n is col
func createTwoDimentionalArray(m, n int) [][]bool {
	result := [][]bool{}
	for i := 0; i < m; i++ {
		result = append(result, make([]bool, n))
	}
	return result
}
