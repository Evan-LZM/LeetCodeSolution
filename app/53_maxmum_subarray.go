package app

//MaxSubArray aaa
func MaxSubArray(nums []int) int {
	if len(nums) == 0 {
		return 0
	}
	result := nums[0]
	sum := 0
	for i := 0; i < len(nums); i++ {
		t := sum + nums[i]
		result = max(result, t)
		if t <= 0 {
			sum = 0
		} else {
			sum = t
		}
	}
	return result
}
