package app

/*SetZeroes Given a m x n matrix, if an element is 0, set its
entire row and column to 0. Do it in-place.

{1, 1, 1},
{1, 0, 1},
{1, 1, 1},

 [1,0,1],
  [0,0,0],
  [1,0,1]
*/

func SetZeroes(matrix [][]int) {
	if len(matrix) == 0 || len(matrix[0]) == 0 {
		return
	}
	row, col := recordZero(matrix)
	for i := 0; i < len(matrix); i++ {
		for j := 0; j < len(matrix[0]); j++ {
			if _, ok := row[i]; ok {
				matrix[i][j] = 0
				continue
			}
			if _, ok := col[j]; ok {
				matrix[i][j] = 0
				continue
			}

		}
	}
}

func recordZero(matrix [][]int) (map[int]bool, map[int]bool) {
	row := make(map[int]bool)
	col := make(map[int]bool)
	for i := 0; i < len(matrix); i++ {
		for j := 0; j < len(matrix[0]); j++ {
			if matrix[i][j] == 0 {
				row[i] = true
				col[j] = true
			}
		}
	}
	return row, col
}
