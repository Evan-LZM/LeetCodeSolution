package main

import (
	"fmt"
	"leetcode/app"
)

/*运算符
假定 A = 60; B = 13; 其二进制数转换为：

A = 0011 1100

B = 0000 1101

-----------------

A&B = 0000 1100

A|B = 0011 1101

A^B = 0011 0001
&	按位与运算符"&"是双目运算符。 其功能是参与运算的两数各对应的二进位相与。	(A & B) 结果为 12, 二进制为 0000 1100
|	按位或运算符"|"是双目运算符。 其功能是参与运算的两数各对应的二进位相或	(A | B) 结果为 61, 二进制为 0011 1101
^	按位异或运算符"^"是双目运算符。 其功能是参与运算的两数各对应的二进位相异或，当两对应的二进位相异时，结果为1。	(A ^ B) 结果为 49, 二进制为 0011 0001
<<	左移运算符"<<"是双目运算符。左移n位就是乘以2的n次方。 其功能把"<<"左边的运算数的各二进位全部左移若干位，由"<<"右边的数指定移动的位数，高位丢弃，低位补0。	A << 2 结果为 240 ，二进制为 1111 0000
>>	右移运算符">>"是双目运算符。右移n位就是除以2的n次方。 其功能是把">>"左边的运算数的各二进位全部右移若干位，">>"右边的数指定移动的位数。
*/

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
	//fmt.Println(app.SolveNQueens2(4))
	// data := [][]int{
	// 	{1, 2, 3},
	// 	{4, 5, 6},
	// 	{7, 8, 9},
	// }
	// app.Rotate(data)
	//fmt.Println(app.LengthOfLastWord("a "))
	//fmt.Println(app.UniquePaths(7, 3))
	// var data = [][]int{
	// 	{0, 0, 0},
	// 	{0, 1, 0},
	// 	{0, 0, 0},
	// }
	// fmt.Println(app.UniquePathsWithObstacles(data))
	//fmt.Println(app.AddBinary("10", "1"))
	//fmt.Println(app.MySqrt(8))
	// var data = [][]int{
	// 	{1, 1, 1},
	// 	{1, 0, 1},
	// 	{1, 1, 1},
	// }
	//	app.SetZeroes(data)
	// var data = [][]int{
	// 	{1},
	// }
	//fmt.Println(app.SearchMatrix(data, 13))
	//fmt.Println(app.Combine(4, 2))
	//fmt.Println(app.Subsets([]int{1, 2, 3}))
	// var data = [][]int{
	// 	{259, 770},
	// 	{448, 54},
	// 	{926, 667},
	// 	{184, 139},
	// 	{840, 118},
	// 	{557, 469},
	// }
	// fmt.Println(app.TwoCitySchedCost(data))
	//fmt.Println(app.ClimbStairs(5))
	// board := [][]byte{
	// 	{'A', 'B', 'C', 'E'},
	// 	{'S', 'F', 'C', 'S'},
	// 	{'A', 'D', 'E', 'E'},
	// }
	// fmt.Println(app.Exist(board, "SEE"))
	// fmt.Println(app.SearchBinary([]int{2, 5, 6, 0, 0, 1, 2}, 0))
	// var root app.TreeNode
	// var right app.TreeNode
	// var left app.TreeNode
	// root.Val = 1
	// right.Val = 2
	// left.Val = 3
	// right.Left = &left
	// root.Right = &right
	// fmt.Println(app.InorderTraversal(&root))
	// data := []byte{'h', 'e', 'l', 'l', 'o'}
	// app.ReverseString(data)
	// insert := [][]int{{1, 2}, {3, 5}, {6, 7}, {8, 10}, {12, 16}}
	// newInterval := []int{4, 8}
	// fmt.Println(app.Insert(insert, newInterval))
	//fmt.Println(app.Coinchange(5, []int{1, 2, 5}))
	//fmt.Println(app.GetPermuation(3, 3))
	//fmt.Println(app.PlusOne([]int{1, 2, 9}))
	// head := app.ListNode{Val: 1}
	// headone := app.ListNode{Val: 2}
	// next1 := app.ListNode{Val: 3}
	// next2 := app.ListNode{Val: 4}
	// next3 := app.ListNode{Val: 5}
	// next2.Next = &next3
	// next1.Next = &next2
	// headone.Next = &next1
	// head.Next = &headone
	// //fmt.Println(app.DeleteDuplicates(&headone))
	// fmt.Println(app.DeleteSortedDuplicates(&headone))
	// for headone.Next != nil {
	// 	fmt.Println(headone.Val)
	// 	headone = *headone.Next
	// }
	// fmt.Println(headone.Val)
	//data := app.RotateRight(&head, 2)
	// data := app.ReverseBetween(&head, 2, 4)
	// for data != nil {
	// 	fmt.Println("result:", data.Val)
	// 	data = data.Next
	// }
	//fmt.Println(app.MinDistance("horse", "ros"))
	//fmt.Println(initialMonthDic(6))
	// now := time.Now()
	// before := time.Now().AddDate(0, 0, -10)
	// fmt.Println("now:", now, ",before:", before)
	// duration := now.Sub(before)
	// fmt.Println("duration:", duration)
	// days := duration.Hours() / 24
	// if days >= 0 && days <= 20 {

	// }
	// fmt.Println("days:", days)

	//app.RemoveDuplicatesSortedArray([]int{1, 1, 1, 2, 3})
	//app.NumDecodings("226")
	//app.RestoreIpAddresses("25525511135")
	//app.NumTrees(3)
	//app.GenerateTrees(3)
	//app.IsInterleave("a", "", "a")
	// root := app.TreeNode{Val: 3}
	// leftone := app.TreeNode{Val: 9}
	// rightone := app.TreeNode{Val: 20}
	// lefttwo := app.TreeNode{Val: 15}
	// righttwo := app.TreeNode{Val: 7}
	// root.Left = &leftone
	// root.Right = &rightone
	// rightone.Left = &lefttwo
	// rightone.Right = &righttwo
	// fmt.Println(app.LevelOrder(&root))
	// s := []string{"hot", "dot", "dog", "lot", "log", "cog"}
	// fmt.Println(app.LadderLength("hit", "cog", s))
	//fmt.Println(app.EvalRPN([]string{"10", "6", "9", "3", "+", "-11", "*", "/", "*", "17", "+", "5", "+"}))
	//fmt.Println(app.SingleNumber([]int{2, 2, 3, 2}))
	//fmt.Println(app.Partition2("aab"))
	//fmt.Println(app.WordBreak("leetcode", []string{"apple", "pen", "leet", "code", "cats", "dog", "sand", "and", "cat"}))
	//fmt.Println(app.ReverseWords("the sky is blue"))
	//app.CompareVersion("1.0.1", "1")
	fmt.Println(app.FindWords([][]byte{{'o', 'a', 'a', 'n'}, {'e', 't', 'a', 'e'}, {'i', 'h', 'k', 'r'}, {'i', 'f', 'l', 'v'}}, []string{"oath", "pea", "eat", "rain"}))
}
