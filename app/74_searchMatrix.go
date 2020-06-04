package app

import "fmt"

/*SearchMatrix Write an efficient algorithm that searches for a value in an m x n matrix.
This matrix has the following properties:
Integers in each row are sorted from left to right.
The first integer of each row is greater than the last integer of the previous row.
*/
func SearchMatrix(matrix [][]int, target int) bool {
	if len(matrix) == 0 || len(matrix[0]) == 0 {
		return false
	}
	//方法1.先转换成一维。再进行二分搜索法。
	data := []int{}
	for i := 0; i < len(matrix); i++ {
		data = append(data, matrix[i]...)
	}
	fmt.Println(data)
	left := 0
	right := len(data) - 1
	for left <= right {
		mid := (left + right) / 2
		if data[mid] == target {
			return true
		} else if data[mid] < target {
			left = mid + 1
		} else {
			right = mid - 1
		}
	}
	return false
}

/*
solution2:use matrix directly
m := len(matrix)
	if m == 0 {
		return false
	}
	n := len(matrix[0])
	if n == 0 {
		return false
	}
	left := 0
	right := m*n - 1
	for left <= right {
		mid := (left + right) / 2
		i := mid / n
		j := mid % n
		if matrix[i][j] == target {
			return true
		} else if matrix[i][j] < target {
			left = mid + 1
		} else {
			right = mid - 1
		}
	}

	return false
*/
