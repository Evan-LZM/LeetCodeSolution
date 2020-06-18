package app

func solve(board [][]byte) {
	if len(board) == 0 {
		return
	}
	if len(board[0]) == 0 {
		return
	}
	//1.找到所有edge边缘的O，2.把o换成Y
	//traversal cols
	for i := 0; i < len(board); i++ {
		helperSolveSurroundedRegion(board, i, 0)
		helperSolveSurroundedRegion(board, i, len(board[0])-1)
	}
	//traversal rows
	for j := 0; j < len(board[0]); j++ {
		helperSolveSurroundedRegion(board, 0, j)
		helperSolveSurroundedRegion(board, len(board)-1, j)
	}
	for i := 0; i < len(board); i++ {
		for j := 0; j < len(board[0]); j++ {
			if board[i][j] == 'O' {
				board[i][j] = 'X'
			} else if board[i][j] == 'Y' {
				board[i][j] = 'O'
			}

		}
	}
}

func helperSolveSurroundedRegion(board [][]byte, i, j int) {
	m := len(board)
	n := len(board[0])
	//i is rows, j is cols
	if i < 0 || j < 0 || i >= m || j >= n || board[i][j] != 'O' {
		return
	}
	board[i][j] = 'Y'
	//find linked nodes
	//1.create directions
	//top:-1,0 right:0,1 bottom:1,0 left:0,-1
	directions := [][]int{{0, 1}, {1, 0}, {-1, 0}, {0, -1}}
	for _, value := range directions {
		ii := value[0] + i
		jj := value[1] + j
		helperSolveSurroundedRegion(board, ii, jj)
	}
}
