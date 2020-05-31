package app

func singleNumber(nums []int) int {
	//^具有记忆性
	r := 0
	for _, n := range nums {
		r ^= n
	}
	return r
}
