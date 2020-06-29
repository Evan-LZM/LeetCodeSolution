package app

/*SingleNumber Given a non-empty array of integers,
every element appears three times except for one,
which appears exactly once. Find that single one.

solution:(d is one)
sum1=3a+3b+3c+d
sum2=a+b+c+d
so:3sum2-sum1=2d
*/
func SingleNumber(nums []int) int {

	//solution1:
	if len(nums) == 0 {
		return 0
	}
	dic := make(map[int]int)
	var sum1, sum2 int
	for _, value := range nums {
		sum1 += value
		dic[value] = 1
	}
	for key := range dic {
		sum2 += key
	}

	return (3*sum2 - sum1) / 2
}

// func SingleNumber2(nums []int) int {
// 	a, b := 0, 0
// 	for _, v := range nums {
// 		a = a ^ v&^b
// 		b = b ^ v&^a
// 	}
// 	return a
// }
