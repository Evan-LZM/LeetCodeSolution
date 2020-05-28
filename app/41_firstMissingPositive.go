package app

func FirstMissingPositive(nums []int) int {
	if len(nums) == 0 {
		return 1
	}
	for i := 0; i < len(nums); i++ {
		for nums[i] != i+1 {
			j := nums[i]
			if j > len(nums) || j < 1 {
				break
			}
			if nums[i] == nums[nums[i]-1] { //nums[i]+1=i
				break
			}
			p := nums[i] - 1
			nums[i], nums[p] = nums[p], nums[i]
		}
		//if nums[i]==i+1: pass
	}
	for i := 0; i < len(nums); i++ {
		if nums[i] != i+1 {
			return i + 1
		}
	}

	return len(nums) + 1
}
