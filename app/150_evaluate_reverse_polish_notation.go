package app

import (
	"strconv"
)

func EvalRPN(tokens []string) int {
	stack := make([]int, len(tokens))
	for _, value := range tokens {
		if value == "+" || value == "/" || value == "-" || value == "*" {
			switch value {
			case "+":
				i := stack[len(stack)-1]
				stack = stack[:len(stack)-1]
				j := stack[len(stack)-1]
				stack = stack[:len(stack)-1]
				m := i + j
				stack = append(stack, m)
			case "-":
				i := stack[len(stack)-1]
				stack = stack[:len(stack)-1]
				j := stack[len(stack)-1]
				stack = stack[:len(stack)-1]
				m := j - i
				stack = append(stack, m)
			case "*":
				i := stack[len(stack)-1]
				stack = stack[:len(stack)-1]
				j := stack[len(stack)-1]
				stack = stack[:len(stack)-1]
				m := j * i
				stack = append(stack, m)
			case "/":
				i := stack[len(stack)-1]
				stack = stack[:len(stack)-1]
				j := stack[len(stack)-1]
				stack = stack[:len(stack)-1]
				m := j / i
				stack = append(stack, m)
			}
		} else {
			i, _ := strconv.Atoi(value)
			stack = append(stack, i)
		}
	}
	return stack[len(stack)-1]
}
