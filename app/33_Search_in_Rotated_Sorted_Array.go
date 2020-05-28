package app

//方法一：放入字典进行检测。但是效率低下
// dic := make(map[int]int)
// 	for index, value := range nums {
// 		dic[value] = index
// 	}
// 	if v, ok := dic[target]; ok {
// 		return v
// 	}
// 	return -1

//方法2:二分搜索技术(尤其是已经排序好的序列)
// if len(nums) == 0 {
// 	return -1
// }
// left, right := 0, len(nums)-1
// for left <= right {
// 	mid := (left + right) / 2
// 	if nums[mid] == target {
// 		return mid
// 	}
// 	if nums[left] <= nums[mid] {
// 		if nums[left] <= target && target <= nums[mid] {
// 			right = mid - 1
// 		} else {
// 			left = mid + 1
// 		}
// 	} else {
// 		if nums[mid] <= target && target <= nums[right] {
// 			left = mid + 1
// 		} else {
// 			right = mid - 1
// 		}

// 	}
// }
// return -1

//Search aaa
func Search(nums []int, target int) int {
	if len(nums) == 0 {
		return -1
	}
	left, right := 0, len(nums)-1
	for left <= right {
		mid := (left + right) / 2
		if nums[mid] == target {
			return mid
		}
		if nums[left] <= nums[mid] {
			if nums[left] <= target && target <= nums[mid] {
				right = mid - 1
			} else {
				left = mid + 1
			}
		} else {
			if nums[mid] <= target && target <= nums[right] {
				left = mid + 1
			} else {
				right = mid - 1
			}

		}
	}
	return -1
}
