package app

import "math"

func minimumTotal(triangle [][]int) int {
	if len(triangle) == 0 {
		return 0
	}
	dp := make([]int, len(triangle))
	dp[0] = triangle[0][0]
	for i := 1; i < len(triangle); i++ {
		for j := i; j >= 0; j-- {
			//处理边界
			if j == i {
				dp[j] = dp[j-1] + triangle[i][j]
			} else if j == 0 {
				dp[j] = dp[0] + triangle[i][j]
			} else {
				dp[j] = min(dp[j], dp[j-1]) + triangle[i][j]
			}
		}
	}
	return getMin(dp)
}

func getMin(arr []int) int {
	res := math.MaxInt32
	for _, value := range arr {
		res = min(value, res)
	}
	return res
}
