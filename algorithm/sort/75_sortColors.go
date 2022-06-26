package main

import "fmt"

//// 快速排序版
//func sortColors(nums []int) {
//	length := len(nums)
//	if length <= 1 {
//		return
//	}
//	var low, high int = 0, length - 1
//	var pivot int = low
//	for low < high {
//		for low < high && nums[low] < nums[pivot] {
//			low++
//		}
//		for low < high && nums[high] > nums[pivot] {
//			high--
//		}
//		nums[low], nums[high] = nums[high], nums[low]
//		high--
//	}
//	sortColors(nums[:low])
//	sortColors(nums[low+1:])
//}

// 统计数字版
func sortColors(nums []int) {
	var cnt0, cnt1 int
	for i := 0; i < len(nums); i++ {
		if nums[i] == 0 {
			cnt0++
		} else if nums[i] == 1 {
			cnt1++
		}
	}
	for i := 0; i < len(nums); i++ {
		if i >= 0 && i < cnt0 {
			nums[i] = 0
		} else if i >= cnt0 && i < cnt0+cnt1 {
			nums[i] = 1
		} else {
			nums[i] = 2
		}
	}
}
func main() {
	d := []int{2, 0, 2, 1, 1, 0}
	sortColors(d)
	fmt.Println(d)
}
