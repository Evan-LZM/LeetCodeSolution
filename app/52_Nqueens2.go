package app

/*SolveNQueens :The n-queens puzzle is the problem of placing n queens on
an n×n chessboard such that no two queens attack each other.
Given an integer n, return all distinct solutions to the n-queens puzzle.

Each solution contains a distinct board configuration of the n-queens' placement,
where 'Q' and '.' both indicate a queen and an empty space respectively.
根据题目的特性，创建一个二维的dp不如创建三个一维dp. 因为二维的需要遍历很多次
因为queen的特效是：横竖，左斜，右斜都是无敌的。所以需要定义三个一维数组，并且遍历横的。就可以解决四种效果
*/
func SolveNQueens2(n int) int {
	solution := [][]string{}
	col := make([]bool, n)
	d1 := make([]bool, 2*n-1)
	d2 := make([]bool, 2*n-1)
	update := 0
	isSolved2(&solution, 0, n, []int{}, col, d1, d2, &update)
	return update
}

//i is the row number and increased by each loop
func isSolved2(result *[][]string, i, n int, temp []int, col, d1, d2 []bool, update *int) {
	if i == n {
		*update++
	}
	//j is the col index
	for j := 0; j < n; j++ {
		if col[j] || d1[i+j] || d2[i-j+n-1] {
			continue
		}
		col[j], d1[i+j], d2[i-j+n-1] = true, true, true
		temp = append(temp, j)
		isSolved2(result, i+1, n, temp, col, d1, d2, update)
		col[j], d1[i+j], d2[i-j+n-1] = false, false, false
		temp = temp[:len(temp)-1]
	}
}
