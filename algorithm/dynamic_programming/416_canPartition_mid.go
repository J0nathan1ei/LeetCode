package main

import "fmt"

func canPartition(nums []int) bool {
	var sum int
	for i := 0; i < len(nums); i++ {
		sum += nums[i]
	}
	if sum%2 == 1 {
		return false
	}
	target := sum / 2
	dp := make([]int, target+1)
	for i := 1; i < target+1; i++ {
		dp[i] = nums[0]
	}
	for i := 1; i < len(nums); i++ {
		for j := target; j >= 0; j-- {
			if dp[j]+nums[i] == target {
				return true
			} else if dp[j]+nums[i] > target {
				dp[j] = dp[j]
			} else {
				dp[j] = max(dp[j], dp[target-j]+nums[i])
			}
		}
	}
	return false
}

func max(x, y int) int {
	if x > y {
		return x
	}
	return y
}

func main() {
	nums := []int{2, 6, 8}
	fmt.Println(canPartition(nums))
}
