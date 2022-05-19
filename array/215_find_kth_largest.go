package main

import "sort"

// 库函数实现
func findKthLargest(nums []int, k int) int {
	length := len(nums)
	if length < k {
		return -1
	}
	sort.Ints(nums)
	return nums[length-k]
}

// 自写的排序算法，见go-demo仓库中的go-sort目录
