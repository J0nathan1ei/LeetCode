package main

import "fmt"

// 解法是标准的动态规划，
// 当前的决策受到过去决策的影响
func minFlipsMonoIncr(s string) int {
	length := len(s)
	if length <= 1 {
		return 0
	}

	dp := make([][]int, length)
	for i := 0; i < length; i++ {
		dp[i] = make([]int, 2)
	}
	// 首字母为0, 1是初始化方式不一样
	if s[0] == '0' {
		dp[0][0], dp[0][1] = 0, 0
	} else if s[0] == '1' {
		dp[0][0], dp[0][1] = 1, 0
	}

	for i := 1; i < length; i++ {
		if s[i] == '0' {
			dp[i][0] = dp[i-1][0]
			dp[i][1] = min(dp[i-1][0], dp[i-1][1]) + 1
		} else {
			dp[i][0] = dp[i-1][0] + 1
			dp[i][1] = min(dp[i-1][0], dp[i-1][1])
		}
	}
	return min(dp[length-1][0], dp[length-1][1])
}

func min(x, y int) int {
	if x < y {
		return x
	}
	return y
}

// 滚动数组法，比较难理解了
func minFlipsMonoIncrII(s string) int {
	length := len(s)
	if length <= 1 {
		return 0
	}

	dp := make([][]int, length)
	for i := 0; i < length; i++ {
		dp[i] = make([]int, 2)
	}

	for i := 1; i < length; i++ {
		if s[i] == '0' {
			dp[i][0] = dp[i-1][0]
			dp[i][1] = min(dp[i-1][0], dp[i-1][1]) + 1
		} else {
			dp[i][0] = dp[i-1][0] + 1
			dp[i][1] = min(dp[i-1][0], dp[i-1][1])
		}
	}
	return min(dp[length-1][0], dp[length-1][1])
}

func main() {
	fmt.Println(minFlipsMonoIncr("100000001010000"))
}
