package main

import (
	"fmt"
	"strconv"
)

func evalRPN(tokens []string) int {
	var numStack []int
	for i := 0; i < len(tokens); i++ {
		if isFirstNum(tokens[i]) {
			num, _ := strconv.Atoi(tokens[i])
			numStack = append(numStack, num)
		} else {
			num1, num2 := numStack[len(numStack)-1], numStack[len(numStack)-2]
			numStack = numStack[:len(numStack)-1]
			switch tokens[i] {
			case "+":
				numStack[len(numStack)-1] = num2 + num1
			case "-":
				numStack[len(numStack)-1] = num2 - num1
			case "*":
				numStack[len(numStack)-1] = num2 * num1
			case "/":
				numStack[len(numStack)-1] = num2 / num1
			}
		}
	}
	return numStack[0]
}
func isFirstNum(s string) bool {
	if (s[0] >= '0' && s[0] <= '9') || (s[0] == '-' && len(s) > 1) {
		return true
	}
	return false
}
func main() {
	d := []string{"10", "6", "9", "3", "+", "-11", "*", "/", "*", "17", "+", "5", "+"}
	fmt.Println(evalRPN(d))
}
