package app

func SolveSudoku(board [][]byte) {
	row, col, blk := [9][9]bool{}, [9][9]bool{}, [9][9]bool{}
	for i := 0; i < 9; i++ {
		row[i] = [9]bool{}
		col[i] = [9]bool{}
		blk[i] = [9]bool{}
	}
	empty := [][]int{}
	for i := 0; i < 9; i++ {
		for j := 0; j < 9; j++ {
			if board[i][j] == '.' {
				empty = append(empty, []int{i, j})
			} else {
				n := board[i][j] - '1'
				row[i][n] = true
				col[j][n] = true
				blk[i/3*3+j/3][n] = true
			}
		}
	}
	if solveSudoHelper(empty, 0, board, row, col, blk) {
		return
	}
}

func solveSudoHelper(empty [][]int, t int, board [][]byte, row, col, blk [9][9]bool) bool {
	if t == len(empty) {
		return true
	}
	i, j := empty[t][0], empty[t][1]
	for n := 1; n <= 9; n++ {
		if row[i][n-1] || col[j][n-1] || blk[i/3*3+j/3][n-1] {
			continue
		}
		row[i][n-1] = true
		col[j][n-1] = true
		blk[i/3*3+j/3][n-1] = true
		board[i][j] = byte(n + '0')
		if solveSudoHelper(empty, t+1, board, row, col, blk) {
			return true
		}
		row[i][n-1] = false
		col[j][n-1] = false
		blk[i/3*3+j/3][n-1] = false
		board[i][j] = '.'
	}
	return false
}
