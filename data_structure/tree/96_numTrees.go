package main

import "fmt"

// 关键是看左右节点的组合情况

func numTrees(n int) int {
	dp := make([]int, n+1)
	dp[0] = 1
	dp[1] = 1
	if n == 1 {
		return dp[n]
	}
	for i := 2; i < n+1; i++ {
		for j := 0; j < i; j++ {
			dp[i] += dp[j] * dp[i-j-1]
		}
	}
	return dp[n]
}

func main() {
	fmt.Println(numTrees(3))
}
