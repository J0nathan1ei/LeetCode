package main

import "fmt"

// 首先想到的是直接模拟操作
// 时间复杂度N^2，空间复杂度N
// time:8min12s
func getModifiedArray(length int, updates [][]int) []int {
	res := make([]int, length)
	for _, nums := range updates {
		for i := nums[0]; i <= nums[1]; i++ {
			res[i] += nums[2]
		}
	}
	return res
}
func main() {
	d := [][]int{
		[]int{2, 4, 6},
		[]int{5, 6, 8},
		[]int{1, 9, -4},
	}
	fmt.Println(getModifiedArray(10, d))
}
