package app

func Coinchange(amount int, coins []int) int {
	dp := make([]int, amount+1)
	dp[0] = 1
	for _, coin := range coins {
		for i := coin; i <= amount; i++ {
			dp[i] += dp[i-coin]
		}
	}
	return dp[amount]
}

//exceed time limit
// func helperCoinChange(amount, sum int, num *int, coins []int, temp []int) {
// 	if sum == amount {
// 		for i := 0; i < len(temp); i++ {
// 			if i+1 <= len(temp)-1 && temp[i+1] > temp[i] {
// 				return
// 			}
// 		}
// 		*num++
// 		return
// 	}
// 	if sum > amount {
// 		return
// 	}
// 	for _, value := range coins {
// 		sum += value
// 		temp = append(temp, value)
// 		helperCoinChange(amount, sum, num, coins, temp)
// 		sum -= value
// 		temp = temp[:len(temp)-1]
// 	}
// }
