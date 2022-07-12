package main

import "fmt"

// 典型的二分法查找范围
// 但是时间复杂度不是纯logN，因为碰见目标元素时，会以logN的时间复杂度来左右扩散
// 时间11min
func searchRange(nums []int, target int) []int {
	res := []int{-1, -1}
	low, high := 0, len(nums)-1
	for low <= high {
		mid := low + (high-low)/2
		if nums[mid] > target {
			high = mid - 1
		} else if nums[mid] < target {
			low = mid + 1
		} else if nums[mid] == target {
			low, high = mid, mid
			for low > 0 && nums[low-1] == target {
				low--
			}
			for high < len(nums)-1 && nums[high+1] == target {
				high++
			}
			res[0] = low
			res[1] = high
			break
		}
	}

	return res
}

// 如果不加遍历，来进行查找的话，可以传入标志变量
func searchRangeII(nums []int, target int) []int {
	res := []int{-1, -1}
	low, high := 0, len(nums)-1
	for low <= high {
		mid := low + (high-low)/2
		if nums[mid] >= target {
			high = mid - 1
		} else if nums[mid] < target {
			low = mid + 1
		} else if nums[mid] == target {
			low, high = mid, mid
			for low > 0 && nums[low-1] == target {
				low--
			}
			for high < len(nums)-1 && nums[high+1] == target {
				high++
			}
			res[0] = low
			res[1] = high
			break
		}
	}

	return res
}

func main() {
	d := []int{}
	fmt.Println(searchRange(d, 6))
}
