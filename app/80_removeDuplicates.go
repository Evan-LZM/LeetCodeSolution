package app

/*RemoveDuplicatesSortedArray Given nums = [1,1,1,2,2,3]
1,1,2,2,3
return 5
Given nums = [0,0,1,1,1,1,2,3,3]
0,0,1,1,2,3,3
return 7
*/

func removeDuplicates(nums []int) int {
	length := 0

	i := 0
	for i < len(nums) {
		p := i
		for i+1 < len(nums) && nums[i] == nums[i+1] {
			i++
		}

		for j := p; j <= i; j++ {
			if j-p+1 <= 2 {
				nums[length] = nums[p]
				length++
			} else {
				break
			}
		}
		i++
	}

	return length
}

// //method:1
// func RemoveDuplicatesSortedArray(nums []int) int {
// 	count := len(nums)
// 	if count <= 2 {
// 		return count
// 	}
// 	replica := 1
// 	for i := 1; i < count; {
// 		if nums[i] == nums[i-1] {
// 			replica++
// 		} else {
// 			replica = 1
// 		}
// 		if replica > 2 {
// 			count--
// 			nums = append(nums[:i], nums[i+1:]...)
// 			replica--
// 		} else {
// 			i++
// 		}
// 	}
// 	return count
// }
