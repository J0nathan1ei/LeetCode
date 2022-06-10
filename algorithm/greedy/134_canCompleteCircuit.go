package main

import (
	"fmt"
)

// // 一开始的思路，计算出path[]，然后选择path[i]最大，且两个和相加大于0的
//func canCompleteCircuit(gas []int, cost []int) int {
//	length := len(gas)
//	path := make([]int, length)
//	var res, sum int
//	for i := 0; i < length; i++ {
//		path[i] = gas[i] - cost[i]
//		sum += path[i]
//	}
//	if sum < 0 {
//		return -1
//	}
//	for i := 0; i < length; i++ {
//		if (i == length-1 && path[i]+path[0] >= 0) ||
//			(i < length-1 && path[i] > 0 && path[i]+path[i+1] >= 0) {
//			res = maxSliceIndex(res, i, path)
//		}
//	}
//	return res
//}
//
//func maxSliceIndex(x, y int, nums []int) int {
//	if nums[x] >= nums[y] {
//		return x
//	}
//	return y
//}

// 答案的贪心法思路，正解要使得局部和大于0，否则更新正解的下标
func canCompleteCircuit(gas []int, cost []int) int {
	length := len(gas)
	path := make([]int, length)
	var res, sum, tmpSum int
	for i := 0; i < length; i++ {
		path[i] = gas[i] - cost[i]
		sum += path[i]
		tmpSum += path[i]
		if tmpSum < 0 {
			res = i + 1
			tmpSum = 0
		}
	}
	if sum < 0 {
		return -1
	}
	return res
}
func main() {
	a := []int{5, 8, 2, 8}
	b := []int{6, 5, 6, 6}
	fmt.Println(canCompleteCircuit(a, b))
}
