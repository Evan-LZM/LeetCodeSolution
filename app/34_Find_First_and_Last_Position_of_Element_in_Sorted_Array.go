package app

//SearchRange aa
func SearchRange(nums []int, target int) []int {
	if len(nums) == 0 {
		return []int{-1, -1}
	}
	return []int{findBinaryLeft(nums, target), findBinaryRight(nums, target)}
}

//找到包含重复排序目标的最左位置
func findBinaryLeft(nums []int, target int) int {
	left, right := 0, len(nums)-1
	for left <= right {
		mid := (left + right) / 2
		if nums[mid] == target {
			right = mid - 1
		} else if nums[mid] < target {
			left = mid + 1
		} else {
			right = mid - 1
		}
	}
	if left >= 0 && len(nums) > left && nums[left] == target {
		return left
	}
	return -1
}

func findBinaryRight(nums []int, target int) int {
	left, right := 0, len(nums)-1
	for left <= right {
		mid := (left + right) / 2
		if nums[mid] == target {
			left = mid + 1
		} else if nums[mid] < target {
			left = mid + 1
		} else {
			right = mid - 1
		}
	}
	if right >= 0 && right < len(nums) && nums[right] == target {
		return right
	}
	return -1
}
