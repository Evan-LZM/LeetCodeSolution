package app

//Jump greedy algo
func Jump(nums []int) int {
	reach, n := 0, len(nums)
	curIndex, count := 0, 0
	for i := 0; i < n; i++ {
		if i > curIndex {
			return -1
		}
		reach = maxDisJump(reach, i+nums[i])
		if i == len(nums)-1 {
			break

		}
		if i == curIndex {
			curIndex = reach
			count++
		}
	}
	return count
}

func maxDisJump(a, b int) int {
	if a > b {
		return a
	}
	return b
}
