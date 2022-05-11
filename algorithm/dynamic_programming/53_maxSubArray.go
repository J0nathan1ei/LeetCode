package main

import (
	"fmt"
)

/*
最大子数组和，典型的动态规划，当前的最大值，可以由过去的最大值推出
*/
func maxSubArray(nums []int) int {
	length := len(nums)
	if length == 1{
		return nums[0]
	}
	var res int = nums[0]
	resArr := make([]int, length)
	resArr[0] = nums[0]
	for i := 1; i < length;i++{
		resArr[i] = max(resArr[i-1]+nums[i], nums[i])
		res = max(res, resArr[i])
	}
	return res
}

func max(x, y int)int{
	if x > y{
		return x
	}
	return y
}

func main(){
	d:= []int{-2,1,-3,4,-1,2,1,-5,4}
	fmt.Println(maxSubArray(d))
}