package main

import (
	"fmt"
)

// 废弃， 因为如果想利用二分法来寻找旋转点的话，要讨论的情况太复杂
//func search(nums []int, target int) bool {
//	var (
//		length         int = len(nums)
//		low, high, mid int = 0, length - 1, 0
//		minIndex       int = 0
//	)
//	// 找到最小值的下标
//	for low <= high {
//		mid = low + (high-low)/2
//		if nums[mid] == target {
//			return true
//		}
//		if nums[mid] < nums[high] {
//			minIndex = mid
//			high = mid - 1
//		} else if nums[mid] < nums[low] {
//			minIndex = mid
//			low = mid + 1
//		}
//	}
//	minIndex = mid
//	// 第一次二分查找
//	low = minIndex
//	for low <= high {
//		mid = low + (high-low)/2
//		if nums[mid] == target {
//			return true
//		} else if nums[mid] < target {
//			low = mid + 1
//		} else if nums[mid] > target {
//			high = mid - 1
//		}
//	}
//	// 第二次二分查找
//	low = 0
//	high = minIndex
//	for low <= high {
//		mid = low + (high-low)/2
//		if nums[mid] == target {
//			return true
//		} else if nums[mid] < target {
//			low = mid + 1
//		} else if nums[mid] > target {
//			high = mid - 1
//		}
//	}
//	return false
//}
func search(nums []int, target int) bool {
	length := len(nums)
	if length == 0 {
		return false
	}
	for i := 0; i < len(nums); i++ {
		if nums[i] == target {
			return true
		}
	}
	return false
}
func main() {
	d := []int{2, 5, 6, 0, 0, 1, 2}
	fmt.Println(search(d, 3))
}
