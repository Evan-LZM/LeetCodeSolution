package app

import "fmt"

/*NumDecodings A message containing letters from A-Z is being
encoded to numbers using the following mapping:
'A' -> 1
'B' -> 2
...
'Z' -> 26
Given a non-empty string containing only digits, determine
the total number of ways to decode it.
solution:

dp[i] stores the value for combination
for example s=226 case
when it comes to 6 index
dp[i]=d[i-1]+dp[i-2]
when s[i] satisfy s[i]>0 for dp[i-1]
when s[i-1:i] satisfy 10<=x<=26
*/
func NumDecodings(s string) int {
	dp := make([]int, len(s)+1)
	dp[0] = 1
	for i := 1; i <= len(s); i++ {
		if i > 1 {
			fmt.Println("i:", i)
			num := 10*int(s[i-2]-'0') + int(s[i-1]-'0')
			if num >= 10 && num <= 26 {
				dp[i] += dp[i-2]
			}
		}
		if int(s[i-1]-'0') > 0 {
			dp[i] += dp[i-1]
		}
		if dp[i] == 0 {
			return 0
		}

	}
	return dp[len(s)]
}
