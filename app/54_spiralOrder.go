package app

/*SpiralOrder 迷宫走法。 给定义一个数组以及方向。当满足某种条件的时候就改变方向
 */
func SpiralOrder(matrix [][]int) []int {
	var result []int
	if len(matrix) == 0 || len(matrix[0]) == 0 {
		return result
	}
	direction := [][]int{
		{0, 1},
		{1, 0},
		{-1, 0},
		{0, -1},
	}
	used := [][]bool{}
	for i := 0; i < len(matrix); i++ {
		used = append(used, make([]bool, len(matrix[0])))
	}
	row, col, d := 0, -1, 0
	for c := 0; c < len(matrix)*len(matrix[0]); {
		xx, yy := row+direction[d][0], col+direction[d][1]
		if xx >= 0 && xx < len(matrix) && yy >= 0 && yy < len(matrix[0]) && !used[xx][yy] {
			used[xx][yy] = true
			result = append(result, matrix[xx][yy])
			row, col = xx, yy
			c++
		} else {
			d = (d + 1) % 4
		}
	}
	return result
}
