package app

func islandPerimeter(grid [][]int) int {
	if len(grid) == 0 || len(grid[0]) == 0 {
		return 0
	}
	res := 0
	for i := 0; i < len(grid); i++ {
		for j := 0; j < len(grid[0]); j++ {
			if grid[i][j] == 0 {
				continue
			}
			res += 4 - getCount(grid, i, j)
		}
	}
	return res
}

func getCount(grid [][]int, i, j int) int {
	count := 0
	for _, direction := range directions {
		ii, jj := i+direction[0], j+direction[1]
		if ii >= 0 && ii < len(grid) && jj >= 0 && jj < len(grid[0]) && grid[ii][jj] == 1 {
			count++
		}
	}
	return count
}
