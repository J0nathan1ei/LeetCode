package main

import "fmt"

// 典型的栈模拟，但是问题主要在于如何处理顺序
// 看了下题解
//
func scoreOfParentheses(s string) int {
	stack := []int{0}
	for _, val := range s {
		if val == '(' {
			stack = append(stack, 0)
		} else {
			pop := stack[len(stack)-1]
			stack = stack[:len(stack)-1]
			stack[len(stack)-1] += max(2*pop, 1)
		}
	}
	return stack[0]
}

func max(x, y int) int {
	if x > y {
		return x
	}
	return y
}

func main() {
	d := "(()())"
	fmt.Println(scoreOfParentheses(d))
}
