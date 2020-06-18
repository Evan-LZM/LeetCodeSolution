package app

func getRow(rowIndex int)[]int{
	res := make([][]int, rowIndex+1)
	res[0] = []int{1}
	if rowIndex == 0 {
		return res[0]
	}
	res[1] = []int{1, 1}
	if rowIndex == 1 {
		return res[1]
	}
	for i := 3; i <= rowIndex+1; i++ {
		temp := make([]int, i)
		temp[0] = 1
		temp[i-1] = 1
		for j := 1; j <= i-2; j++ {
			temp[j] = res[i-2][j] + res[i-2][j-1]
		}
		res[i-1] = temp
	}
	return res[rowIndex]
}