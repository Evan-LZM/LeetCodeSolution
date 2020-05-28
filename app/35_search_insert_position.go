package app

import "fmt"

//SearchInsert aaa
//int{1, 3, 5, 6}, 4
//    0, 1, 2, 3
func SearchInsert(nums []int, target int) int {
	left, right := 0, len(nums)-1
	pos := -1
	for left <= right {
		mid := (left + right) / 2
		fmt.Println("mid:", mid)
		if nums[mid] == target {
			pos = mid
			break
		} else if nums[mid] <= target {
			fmt.Println("before_l:", left)
			left = mid + 1
			fmt.Println("after:", left)
		} else {
			fmt.Println("before_r:", right)
			right = mid - 1
			fmt.Println("after:", right)
		}
	}
	if pos != -1 {
		return pos
	}
	fmt.Println("pos:", pos, left, right)
	return left
}
