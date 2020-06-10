package app

//SearchInsert aaa
//int{1, 3, 5, 6}, 4
//    0, 1, 2, 3
func SearchInsert(nums []int, target int) int {
	left, right := 0, len(nums)-1
	pos := -1
	for left <= right {
		mid := (left + right) / 2
		if nums[mid] == target {
			pos = mid
			break
		} else if nums[mid] <= target {
			left = mid + 1
		} else {
			right = mid - 1
		}
	}
	if pos != -1 {
		return pos
	}
	return left
}
