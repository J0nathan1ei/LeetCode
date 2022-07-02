package main

import "fmt"

// 基本解法，遍历两个数组，时间复杂度N
func game(guess []int, answer []int) int {
	var res int
	for i := 0; i < len(guess); i++ {
		if guess[i] == answer[i] {
			res++
		}
	}
	return res
}

func main() {
	fmt.Println()
}
