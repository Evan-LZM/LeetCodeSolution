package app

func generate(numRows int) [][]int {
	res := make([][]int, numRows)
	if numRows == 0 {
		return res
	}
	res[0] = []int{1}
	if numRows == 1 {
		return res
	}
	res[1] = []int{1, 1}
	if numRows == 2 {
		return res
	}
	for i := 3; i <= numRows; i++ {
		temp := make([]int, i)
		temp[0] = 1
		temp[i-1] = 1
		for j := 1; j <= i-2; j++ {
			temp[j] = res[i-2][j] + res[i-2][j-1]
		}
		res[i-1] = temp
	}
	return res
}
