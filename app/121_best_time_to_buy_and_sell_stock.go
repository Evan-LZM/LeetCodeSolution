package app

func maxProfit(prices []int) int {
	if len(prices) <= 1 {
		return 0
	}
	var temp, res int
	for i := 1; i < len(prices); i++ {
		next := temp + prices[i] - prices[i-1]
		res = max(res, next)
		if next > 0 {
			temp = next
		} else {
			temp = 0
		}
	}
	return res
}
