package app

import "fmt"

//IsValidSudoku aa
func IsValidSudoku(board [][]byte) bool {
	col := makeTwoDimentionalArray()
	row := makeTwoDimentionalArray()
	blk := makeTwoDimentionalArray()
	for i := 0; i < 9; i++ {
		for j := 0; j < 9; j++ {
			if board[i][j] == '.' {
				continue
			}
			n := board[i][j] - '1'
			if col[j][n] {
				return false
			}
			if row[i][n] {
				return false
			}
			b := j/3*3 + i/3
			if blk[b][n] {
				return false
			}
			col[j][n] = true
			row[i][n] = true
			blk[b][n] = true
		}
	}
	fmt.Println("col:")
	fmt.Println(col)
	fmt.Println("row:")
	fmt.Println(row)
	fmt.Println("blk:")
	fmt.Println(blk)
	return true
}

func makeTwoDimentionalArray() [][]bool {
	var twoDim [][]bool
	for i := 0; i < 9; i++ {
		oneD := make([]bool, 9)
		twoDim = append(twoDim, oneD)
	}
	return twoDim
}
