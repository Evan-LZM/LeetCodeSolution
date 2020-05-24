package app

import (
	"math"
	"sort"
)

//NextPermutation cc
func NextPermutation(nums []int) {
	if len(nums) <= 1 {
		return
	}
	i := len(nums) - 1
	for i >= 0 {
		if i-1 >= 0 && nums[i-1] >= nums[i] {
			i = i - 1
		} else {
			break
		}
	}
	if i == 0 {
		sort.Ints(nums)
		return
	}
	n := nums[i-1]
	p := -1
	mn := math.MaxInt32
	for j := i; j < len(nums); j++ {
		if nums[j] > n {
			if nums[j] < mn {
				mn = nums[j]
				p = j
			}
		}
	}
	nums[i-1], nums[p] = nums[p], nums[i-1]
	sort.Ints(nums[i:])
	return
}
