package app

func GenerateMatrix(n int) [][]int {
	var data [][]int

	direction := [][]int{
		{0, 1},
		{1, 0},
		{0, -1},
		{-1, 0},
	}
	var used [][]bool
	for i := 0; i < n; i++ {
		used = append(used, make([]bool, n))
		data = append(data, make([]int, n))
	}
	row, col, deep := 0, -1, 0
	for c := 0; c < n*n; {
		xx, yy := row+direction[deep][0], col+direction[deep][1]
		if xx >= 0 && yy >= 0 && xx < n && yy < n && !used[xx][yy] {
			used[xx][yy] = true
			data[xx][yy] = c + 1
			row, col = xx, yy
			c++
		} else {
			deep = (1 + deep) % 4
		}

	}
	return data
}
