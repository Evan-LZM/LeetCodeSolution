package app

//PermuteUnique Given a collection of numbers that might contain duplicates, return all possible unique permutations.
//1.recursion
func PermuteUnique(nums []int) [][]int {
	var result [][]int
	used := make([]bool, len(nums))
	temp := []int{}
	helpPermute(nums, used, temp, &result)
	return result
}

func helpPermute(nums []int, used []bool, temp []int, result *[][]int) {
	if len(nums) == len(temp) {
		r := append([]int{}, temp...)
		*result = append(*result, r)
	}
	for i := 0; i < len(nums); i++ {
		if used[i] {
			continue
		}
		used[i] = true
		temp = append(temp, nums[i])
		helpPermute(nums, used, temp, result)
		used[i] = false
		temp = temp[:len(temp)-1]
	}
}
