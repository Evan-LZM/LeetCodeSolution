package app

/*
131
151
421

You can only move either down or right at any point in time.
*/

func MinPathSum(grid [][]int) int {
	if len(grid) == 0 {
		return 0
	}
	dp := createDP(len(grid), len(grid[0]))
	//initial dp
	//dp[0][0] = grid[0][0]
	for i := 0; i < len(grid); i++ {
		for j := 0; j < len(grid[0]); j++ {
			if i == 0 && j == 0 {
				dp[i][j] = grid[i][j]
				continue
			}
			if i == 0 && j != 0 {
				dp[i][j] = dp[i][j-1] + grid[i][j]
				continue
			}
			if i > 0 && j == 0 {
				dp[i][j] = dp[i-1][j] + grid[i][j]
				continue
			}
			dp[i][j] = minReturn(dp[i-1][j], dp[i][j-1]) + grid[i][j]
		}
	}
	return dp[len(grid)-1][len(grid[0])-1]
}

//m=rows,n=col
func createDP(m, n int) [][]int {
	var data [][]int
	for i := 0; i < m; i++ {
		data = append(data, make([]int, n))
	}
	return data
}

func minReturn(a, b int) int {
	if a > b {
		return b
	}
	return a
}
