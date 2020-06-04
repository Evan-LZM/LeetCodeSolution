package app

/*Combine Given two integers n and k, return all possible
combinations of k numbers out of 1 ... n.
Input: n = 4, k = 2
Output:
[
  [2,4],
  [3,4],
  [2,3],
  [1,2],
  [1,3],
  [1,4],
]
*/
func Combine(n int, k int) [][]int {
	if n <= 0 {
		return [][]int{}
	}
	if n == 1 {
		return [][]int{{1}}
	}
	data := [][]int{}
	used := make([]bool, n+1)
	for i := 0; i <= n; i++ {
		used[i] = false
	}
	helperCombine(&data, 1, n, k, []int{}, used)
	return data
}

func helperCombine(result *[][]int, i, n, k int, temp []int, used []bool) {
	if len(temp) == k {
		r := append([]int{}, temp...)
		*result = append(*result, r)
		return
	}
	for i <= n {
		if used[i] {
			i++
		} else {
			temp = append(temp, i)
			used[i] = true
			helperCombine(result, i+1, n, k, temp, used)
			used[i] = false
			temp = temp[:len(temp)-1]
			i++
		}
	}
}
