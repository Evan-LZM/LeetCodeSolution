package app

import (
	"sort"
)

func FourSum(nums []int, target int) [][]int {
	var result [][]int
	if len(nums) < 4 {
		return result
	}
	sort.Ints(nums)
	for i := 0; i < len(nums); i++ {
		if i > 0 && nums[i] == nums[i-1] {
			continue
		}
		for j := i + 1; j < len(nums); j++ {
			if j > i+1 && nums[j] == nums[j-1] {
				continue
			}
			k := j + 1
			l := len(nums) - 1
			for k < l {
				sum := nums[i] + nums[j] + nums[k] + nums[l]
				if sum == target {
					result = append(result, []int{nums[i], nums[j], nums[k], nums[l]})
					k++
					for ; k < l && nums[k] == nums[k-1]; k++ {
					}
				} else if sum < target {
					k++
				} else {
					l--
				}
			}
		}

	}
	return result
}

func FourSumV2(nums []int, target int) [][]int {
	var result [][]int
	if len(nums) < 4 {
		return result
	}
	sort.Ints(nums)
	for i := 0; i < len(nums)-3; i++ {
		if i > 0 && nums[i] == nums[i-1] {
			continue
		}
		for j := i + 1; j < len(nums)-2; j++ {
			if j > i+1 && nums[j] == nums[j-1] {
				continue
			}
			k := j + 1
			l := len(nums) - 1
			for k < l {
				sum := nums[i] + nums[j] + nums[k] + nums[l]
				if sum == target {
					result = append(result, []int{nums[i], nums[j], nums[k], nums[l]})
					k++
					l--
					for k < l && nums[k] == nums[k-1] {
						k++
					}
					for k < l && nums[l] == nums[l+1] {
						l--
					}

				} else if sum < target {
					k++
				} else {
					l--
				}
			}
		}

	}
	return result
}
