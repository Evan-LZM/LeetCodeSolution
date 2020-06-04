package app

/*UniquePathsWithObstacles A robot is located at the top-left corner of a m x n grid
(marked 'Start' in the diagram below).
The robot can only move either down or right at any point in time.
The robot is trying to reach the bottom-right corner of the grid
(marked 'Finish' in the diagram below).
Now consider if some obstacles are added to the grids.
How many unique paths would there be?

尝试使用一维的dp来解决。
*/
func UniquePathsWithObstacles(obstacleGrid [][]int) int {
	m := len(obstacleGrid) //row
	if m <= 0 {
		return 0
	}
	n := len(obstacleGrid[0]) //col
	if n <= 0 {
		return 0
	}
	dp := make([]int, n)
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if obstacleGrid[i][j] == 1 {
				dp[j] = 0
				continue
			}
			if i == 0 && j == 0 {
				dp[j] = 1
				continue
			}
			up, left := 0, 0
			if i > 0 && obstacleGrid[i-1][j] != 1 {
				up = dp[j]
			}
			if j > 0 && obstacleGrid[i][j-1] != 1 {
				left = dp[j-1]
			}
			dp[j] = up + left
		}
	}
	return dp[n-1]
}
