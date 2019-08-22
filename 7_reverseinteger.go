package leetcodesolution

func reverse(x int) int {
	var Maxint int32 = 2147483647
	if x > int(Maxint) || x < -int(Maxint) {
		return 0
	}
	temp := 0
	run := x
	if x < 0 {
		run = -run
	}
	for run > 0 {
		temp *= 10
		digit := run % 10
		temp += digit
		run /= 10
	}
	if temp > int(Maxint) {
		return 0
	}
	if x < 0 {
		return -temp
	}
	return temp
}
