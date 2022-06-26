package main

import "fmt"

//// 尝试哈希表，22min没有想出来
//func longestConsecutive(nums []int) int {
//	var res int
//	length := len(nums)
//	m := make(map[int]int, length)
//	for i := 0; i < length; i++ {
//		m[nums[i]] = 1
//	}
//	for i := 0; i < length; i++ {
//		if _, ok := m[nums[i]-1]; ok {
//			m[nums[i]]++
//			m[nums[i]-1]++
//		}
//		if _, ok := m[nums[i]+1]; ok {
//			m[nums[i]]++
//			m[nums[i]+1]++
//		}
//		res = max(res, m[nums[i]])
//	}
//	return res
//}
//
//func max(x, y int) int {
//	if x > y {
//		return x
//	}
//	return y
//}

// 题解的方法，理解后重写一遍
func longestConsecutive(nums []int) int {
	var res int
	m := map[int]bool{}
	for _, v := range nums {
		m[v] = true
	}
	for _, v := range nums {
		// 这一步判断之前的元素是否存在，成功避免了n^2的时间复杂度
		// 使得同一序列只有一次顺序查找
		if !m[v-1] {
			tmpNum := 1
			for m[v+1] {
				v++
				tmpNum++
			}
			if res < tmpNum {
				res = tmpNum
			}
		}
	}
	return res
}

func main() {
	d := []int{0, 3, 7, 2, 5, 8, 4, 6, 0, 1}
	fmt.Println(longestConsecutive(d))
}
