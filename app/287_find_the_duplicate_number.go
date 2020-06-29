package app

func findDuplicate(nums []int) int {
	slow := 0
	fast := 0
	//1.判断是否是闭环
	for true {
		fast = nums[nums[fast]]
		slow = nums[slow]
		if slow == fast {
			break
		}
	}
	entry := 0
	for entry != slow {
		entry = nums[entry]
		slow = nums[slow]
	}
	return entry
}
