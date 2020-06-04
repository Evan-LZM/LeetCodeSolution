package app

/*UniquePaths
A robot is located at the top-left corner of a m x n grid (marked 'Start' in the diagram below).

The robot can only move either down or right at any point in time. The robot is trying to reach the bottom-right corner of the grid (marked 'Finish' in the diagram below).
*/
func UniquePaths(m int, n int) int {
	if m == 0 && n == 0 {
		return 0
	}
	if m == 1 || n == 1 {
		return 1
	}
	dp := initialDP(n, m)
	dp[0][0] = 1
	for i := 0; i < n; i++ {
		for j := 0; j < m; j++ {
			if i == 0 && j-1 >= 0 {
				dp[i][j] = 1
			} else if j == 0 && i-1 >= 0 {
				dp[i][j] = 1
			} else if i-1 >= 0 && j-1 >= 0 {
				dp[i][j] = dp[i][j-1] + dp[i-1][j]
			}
		}
	}
	return dp[n-1][m-1]
}

//m is row, n is col
func initialDP(m int, n int) [][]int {
	dp := [][]int{}
	for i := 0; i < m; i++ {
		dp = append(dp, make([]int, n))
	}
	return dp
}

/*
Better Solution
	if m<=0||n<=0{
		return 0
	}
	dp:=make([]int,n)

*/
