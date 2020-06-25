package app

func maxProfit2(prices []int) int {
	if len(prices) <= 1 {
		return 0
	}
	var temp, res int
	for i := 1; i < len(prices); i++ {
		temp = prices[i] - prices[i-1]
		if temp > 0 {
			res += temp
		}
	}
	return res
}
