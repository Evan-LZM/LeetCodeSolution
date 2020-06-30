package app

func maxProduct(nums []int) int {

	mx, mn := nums[0], nums[0]
	res := nums[0]
	for i := 1; i < len(nums); i++ {
		mx = max(max(mx*nums[i], mn*nums[i]), nums[i])
		mn = min(min(mx*nums[i], mn*nums[i]), nums[i])
		res = max(res, mx)
	}
	return res
}
