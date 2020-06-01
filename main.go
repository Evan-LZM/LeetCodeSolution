package main

import (
	"fmt"
	"leetcode/app"
)

func main() {
	fmt.Println("print--result---")
	//app.Convert('leetcode', 3)
	//fmt.Println(app.IsPalindromeV3(121))
	//fmt.Println(app.MaxArea([]int{1, 8, 6, 2, 5, 4, 8, 3, 7}))
	//	fmt.Println(app.IntToRoman(10))
	//fmt.Println(app.RomanToInt('X'))
	//	fmt.Println(app.LongestCommonPrefix([]string{'flower', 'flow', 'flight'}))
	//fmt.Println(app.ThreeSumClosest([]int{1, 1, 1, 0}, 100))
	//fmt.Println(app.LetterCombinationsRecursion('234'))
	//fmt.Println(app.FourSumV2([]int{1, 0, -1, 0, -2, 2}, 0))

	// tail := app.ListNode{
	// 	Val:  5,
	// 	Next: nil,
	// }
	// num4 := app.ListNode{
	// 	Val:  4,
	// 	Next: &tail,
	// }
	// num3 := app.ListNode{
	// 	Val:  3,
	// 	Next: &num4,
	// }
	// num2 := app.ListNode{
	// 	Val:  2,
	// 	Next: &num3,
	// }
	// head := app.ListNode{
	// 	Val:  1,
	// 	Next: &num2,
	// }
	// fmt.Println(app.RemoveNthFromEnd(&head, 2))
	//fmt.Println(app.IsValid('()'))
	//fmt.Println(app.GenerateParenthesis(3))
	//fmt.Println(app.RemoveDuplicates([]int{1, 1, 2}))
	//app.Divide(4, 4)
	//fmt.Println(app.FindSubstring('wordgoodgoodgoodbestword', []string{'word', 'good', 'best', 'word'}))
	//app.NextPermutation([]int{1, 2, 3})
	//fmt.Println(app.LongestValidParentheses(')()())'))
	//fmt.Println(app.Search([]int{4, 5, 6, 7, 0, 1, 2}, 0))
	//	fmt.Println(app.SearchRange([]int{1}, 1))
	//fmt.Println(app.SearchInsert([]int{1, 3, 5, 6}, 4))
	// board := [][]byte{{'5', '3', '.', '.', '7', '.', '.', '.', '.'},
	// 	{'6', '.', '.', '1', '9', '5', '.', '.', '.'},
	// 	{'.', '9', '8', '.', '.', '.', '.', '6', '.'},
	// 	{'8', '.', '.', '.', '6', '.', '.', '.', '3'},
	// 	{'4', '.', '.', '8', '.', '3', '.', '.', '1'},
	// 	{'7', '.', '.', '.', '2', '.', '.', '.', '6'},
	// 	{'.', '6', '.', '.', '.', '.', '2', '8', '.'},
	// 	{'.', '.', '.', '4', '1', '9', '.', '.', '5'},
	// 	{'.', '.', '.', '.', '8', '.', '.', '7', '9'}}
	//fmt.Println(app.IsValidSudoku(board))
	//app.SolveSudoku(board)
	//fmt.Println(app.CombinationSum([]int{1, 2, 3, 5, 7, 8}, 8))
	//fmt.Println(app.CombinationSum2([]int{10, 1, 2, 7, 6, 1, 5}, 8))
	//fmt.Println(app.FirstMissingPositive([]int{2, 2, 2, 2}))
	//fmt.Println(app.Trap([]int{0, 1, 0, 2, 1, 0, 1, 3, 2, 1, 2, 1}))
	//fmt.Println(app.Multiply("70", "80"))
	//fmt.Println(app.IsMatch("aasdfasdfasdfasdfas", "aasdf.*asdf.*asdf.*asdf.*s"))
	//fmt.Println(app.PermuteUnique2([]int{1, 1, 2}))
	//fmt.Println(app.GroupAnagrams([]string{"eat", "tea", "tan", "ate", "nat", "bat"}))
	//	fmt.Println(app.MyPow(2.1, 3))
	//fmt.Println(app.MaxSubArray([]int{-2, 1, -3, 4, -1, 2, 1, -5, 4}))
	// data := [][]int{
	// 	{1, 2, 5},
	// 	{3, 2, 1},
	// }
	// fmt.Println(app.MinPathSum(data))
	//fmt.Println(app.CanJump([]int{2, 3, 1, 1, 4}))
	//fmt.Println(app.Jump([]int{2, 3, 1, 1, 4}))
	// data := [][]int{{1, 2, 3}, {4, 5, 6}, {7, 8, 9}}
	// fmt.Println(app.SpiralOrder(data))
	//fmt.Println(app.GenerateMatrix(3))
	// data := [][]int{
	// 	{1, 3},
	// 	{2, 6},
	// 	{8, 10},
	// 	{15, 18},
	// }
	// fmt.Println(app.Merge(data))
	fmt.Println(app.SolveNQueens2(4))
}
