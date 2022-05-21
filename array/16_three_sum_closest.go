package main

import (
	"fmt"
	"math"
	"sort"
)

/*一开始没有想出什么办法，看了下答案的想法，手动又实现了一遍 排序+双指针
注意一开始res的最大值选取，第一次选择MaxInt，结果导致加数字后溢出了
*/
func threeSumClosest(nums []int, target int) int {
	var res int = math.MaxInt32 // 这里一开始写的是MaxInt，遇到负数时发生溢出了
	length := len(nums)

	sort.Ints(nums)
	for i := 0; i < length; i++ {
		cur := nums[i]
		left, right := i+1, length-1
		for left < right {
			sum := cur + nums[left] + nums[right]
			res = minAbs(res, sum, target)
			if sum < target {
				left++
			} else if sum > target {
				right--
			} else {
				return sum
			}
		}

	}
	return res
}

func minAbs(x, y, target int) int {
	if abs(x, target) < abs(y, target) {
		return x
	}
	return y
}

func abs(x, y int) int {
	if x > y {
		res := x - y
		return res
	}
	return y - x
}

func main() {
	d := []int{1, 1, 1, 1}
	fmt.Println(threeSumClosest(d, -100))
}
