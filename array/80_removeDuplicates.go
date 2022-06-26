package main

import "fmt"

//// 数组截取版
//func removeDuplicates(nums []int) int {
//	length := len(nums)
//	if length <= 2 {
//		return length
//	}
//	var cnt int
//	for i := 1; i < length; i++ {
//		if nums[i] == nums[i-1] {
//			if cnt == 2 {
//				nums = append(nums[:i], nums[i+1:]...)
//				length--
//				// 原地删除元素的时候，后面的元素前移，因此i必须减1
//				i--
//			} else {
//				cnt = 2
//			}
//		} else {
//			cnt = 0
//		}
//	}
//	return length
//}

// 双指针移动版
func removeDuplicates(nums []int) int {
	length := len(nums)
	if length <= 2 {
		return length
	}
	low, high := 2, 2
	for high < length {
		if nums[low-2] != nums[high] {
			nums[low] = nums[high]
			low++
		}
		high++
	}
	return low
}
func main() {
	d := []int{0, 0, 1, 1, 1, 1, 2, 3, 3}
	fmt.Println(removeDuplicates(d))
}
