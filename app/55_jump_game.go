package app

func CanJump(nums []int) bool {
	n, reach := len(nums), 0
	for i := 0; i < n; i++ {
		if i > reach || reach >= n-1 {
			break
		}
		reach = maxDis(reach, i+nums[i])
	}
	return reach > n-1
}
func maxDis(a, b int) int {
	if a > b {
		return a
	}
	return b
}
