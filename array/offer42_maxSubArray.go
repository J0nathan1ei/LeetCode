package main

import "fmt"

func maxSubArray(nums []int) int {
	var res int
	length := len(nums)
	dp := make([]int, length)
	dp[0] = nums[0]
	res = nums[0]
	for i := 1; i < length; i++ {
		dp[i] = max(nums[i], nums[i]+dp[i-1])
		res = max(res, dp[i])
	}
	return res
}

func main() {
	d := []int{-2, 1, -3, 4, -1, 2, 1, -5, 4}
	fmt.Println(maxSubArray(d))
}
