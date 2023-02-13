package main

import "fmt"

func minCostClimbingStairs(cost []int) int {
	var length int = len(cost)
	if length < 3 {
		return min(cost[0], cost[1])
	}

	dp := make([]int, length+1)
	var i int
	for i = 2; i <= length; i++ {
		dp[i] = minVal(dp, cost, i-1, i-2)
	}
	return dp[length]
}

func minVal(dp, nums []int, i, j int) int {
	if dp[i]+nums[i] < dp[j]+nums[j] {
		return dp[i] + nums[i]
	} else {
		return dp[j] + nums[j]
	}
}

func min(x, y int) int {
	if x < y {
		return x
	}
	return y
}

func main() {
	cost := []int{1, 100, 1, 1, 1, 100, 1, 1, 100, 1}
	d := minCostClimbingStairs(cost)
	fmt.Println(d)
}
