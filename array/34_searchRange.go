package main

import "fmt"

func searchRange(nums []int, target int) []int {
	length := len(nums)
	res := []int{-1, -1}
	if length == 0 {
		return res
	} else if target < nums[0] || target > nums[length-1] {
		return res
	}
	low, high, mid := 0, length-1, length/2-1
	for low < high {

	}

}

func main() {
	fmt.Println()
}
+