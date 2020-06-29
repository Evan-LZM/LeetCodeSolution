package app

func WordBreak(s string, wordDict []string) bool {
	//using dp
	dp := make([]bool, len(s)+1)
	dp[0] = true
	for i := 1; i <= len(s); i++ {
		for _, w := range wordDict {
			j := i - len(w)
			if j >= 0 && dp[j] && s[j:i] == w {
				dp[i] = true
				break
			}
		}
	}
	return dp[len(s)]
}
