package app

var directions = [][]int{
	{0, 1},
	{0, -1},
	{1, 0},
	{-1, 0},
}

func Exist(board [][]byte, word string) bool {
	used := make([][]bool, len(board))
	for i := 0; i < len(board); i++ {
		used[i] = make([]bool, len(board[i]))
	}
	for i := 0; i < len(board); i++ {
		for j := 0; j < len(board[i]); j++ {
			if dfs(board, used, word, i, j) {
				return true
			}
		}
	}
	return false
}

func dfs(board [][]byte, visited [][]bool, word string, i, j int) bool {
	if len(word) == 0 {
		return true
	}
	if i < 0 || j < 0 || i >= len(board) || j >= len(board[0]) || visited[i][j] || board[i][j] != word[0] {
		return false
	}
	visited[i][j] = true
	for _, d := range directions {
		if dfs(board, visited, word[1:], i+d[0], j+d[1]) {
			return true
		}
	}
	visited[i][j] = false
	return false
}
