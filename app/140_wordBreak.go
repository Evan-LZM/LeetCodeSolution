package app

//WordBreak2 return all answer
func WordBreak2(s string, wordDict []string) []string {
	var void struct{}
	dict := make(map[string]struct{})
	for _, w := range wordDict {
		if len(w) == 0 {
			continue
		}
		dict[w] = void
	}
	if !canWordBreak(s, dict) {
		return []string{}
	}
	return wordBreakHelper(s, dict)
}

func canWordBreak(s string, dict map[string]struct{}) bool {
	dp := make([]bool, len(s)+1)
	dp[0] = true
	for i := 1; i <= len(s); i++ {
		for w := range dict {
			j := i - len(w)
			if len(w) > i || !dp[j] || w != s[j:i] {
				continue
			}
			dp[i] = true
			break
		}
	}
	return dp[len(s)]
}

func wordBreakHelper(s string, dict map[string]struct{}) []string {
	dp := make([][]string, len(s)+1)
	dp[0] = []string{""}
	for i := 1; i < len(s); i++ {
		for w := range dict {
			if len(w) > i || len(dp[i-len(w)]) == 0 || w != s[i-len(w):i] {
				continue
			}
			for _, pw := range dp[i-len(w)] {
				if len(pw) == 0 {
					dp[i] = append(dp[i], w)
				} else {
					dp[i] = append(dp[i], pw+" "+w)
				}
			}
		}
	}
	return dp[len(s)]
}

/*solution2:

func wordBreak(s string, wordDict []string) []string {
	words := make(map[string]bool)
	for _, w := range wordDict {
		words[w] = true
	}
	return helperWordBreak(s, words, make(map[int][]string), 0)
}

func helperWordBreak(s string, wordDict map[string]bool, record map[int][]string, pos int) []string {
	if res, ok := record[pos]; ok {
		return res
	}
	res := []string{}
	for i := pos + 1; i <= len(s); i++ {
		if wordDict[s[pos:i]] {
			if i != len(s) {
				rest := helperWordBreak(s, wordDict, record, i)
				for _, r := range rest {
					res = append(res, s[pos:i]+" "+r)
				}
			} else {
				res = append(res, s[pos:i])
				break
			}
		}
	}
	record[pos] = res
	return res
}
*/
