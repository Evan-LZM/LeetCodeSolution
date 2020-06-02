package app

import "fmt"

/*
Rotate :You are given an n x n 2D matrix representing an image.
Rotate the image by 90 degrees (clockwise).

Note:
You have to rotate the image in-place, which means you have to modify the input
2D matrix directly. D
O NOT allocate another 2D matrix and do the rotation.
*/
func Rotate(matrix [][]int) {
	if len(matrix) == 1 {
		return
	}
	n := len(matrix)
	mid := (n - 1) / 2
	if n%2 == 0 {
		for i := 0; i <= mid; i++ {
			for j := 0; j < n; j++ {
				SwapRotate(&matrix[i][j], &matrix[n-i-1][j])
			}
		}
	} else {
		for i := 0; i < mid; i++ {
			for j := 0; j < n; j++ {
				SwapRotate(&matrix[i][j], &matrix[n-i-1][j])
			}
		}
	}
	for i := 0; i < n; i++ {
		for j := 0; j < i; j++ {
			SwapRotate(&matrix[i][j], &matrix[j][i])
		}
	}
	fmt.Println(matrix)
}

func SwapRotate(a, b *int) {
	*a, *b = *b, *a
}