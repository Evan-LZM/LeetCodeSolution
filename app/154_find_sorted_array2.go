package app

func findMin2(nums []int) int {
	l, r := 0, len(nums)-1
	for l < r {
		mid := (l + r) / 2
		if nums[l] > nums[mid] {
			r = mid
		} else if nums[mid] > nums[r] {
			l = mid + 1
		} else if nums[mid] == nums[r] {
			r = r - 1
		} else {
			l = l - 1
		}
	}
	return nums[l]
}
