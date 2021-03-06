package app

import "strings"

/*SolveNQueens :The n-queens puzzle is the problem of placing n queens on
an n×n chessboard such that no two queens attack each other.
Given an integer n, return all distinct solutions to the n-queens puzzle.

Each solution contains a distinct board configuration of the n-queens' placement,
where 'Q' and '.' both indicate a queen and an empty space respectively.
根据题目的特性，创建一个二维的dp不如创建三个一维dp. 因为二维的需要遍历很多次
因为queen的特效是：横竖，左斜，右斜都是无敌的。所以需要定义三个一维数组，并且遍历横的。就可以解决四种效果
*/
func SolveNQueens(n int) [][]string {
	solution := [][]string{}
	col := make([]bool, n)
	d1 := make([]bool, 2*n-1)
	d2 := make([]bool, 2*n-1)
	isSolved(&solution, 0, n, []int{}, col, d1, d2)
	return solution
}

//i is the row number and increased by each loop
func isSolved(result *[][]string, i, n int, temp []int, col, d1, d2 []bool) {
	if i == n {
		//temp store the result based on index
		var data []string
		for m := 0; m < len(temp); m++ {
			data = append(data, createString(temp[m], n))
		}
		*result = append(*result, data)
	}
	//j is the col index
	for j := 0; j < n; j++ {
		if col[j] || d1[i+j] || d2[i-j+n-1] {
			continue
		}
		col[j], d1[i+j], d2[i-j+n-1] = true, true, true
		temp = append(temp, j)
		isSolved(result, i+1, n, temp, col, d1, d2)
		col[j], d1[i+j], d2[i-j+n-1] = false, false, false
		temp = temp[:len(temp)-1]
	}
}

//c is the index 0<=c<n
func createString(c, n int) string {
	var sub strings.Builder
	for i := 0; i < c; i++ {
		sub.WriteByte('.')
	}
	sub.WriteByte('Q')
	for i := c + 1; i < n; i++ {
		sub.WriteByte('.')
	}
	return sub.String()
}
