package main

import "fmt"

func exchange(nums []int) []int {
	low, high := 0, len(nums)-1
	for low < high {
		for low < high && nums[low]%2 == 1 {
			low++
		}
		for low < high && nums[high]%2 == 0 {
			high--
		}
		nums[low], nums[high] = nums[high], nums[low]
	}
	return nums
}
func exchangeByFunc(nums []int, cmp func(x int) bool) []int {
	low, high := 0, len(nums)-1
	for low < high {
		for low < high && cmp(nums[low]) {
			low++
		}
		for low < high && !cmp(nums[high]) {
			high--
		}
		nums[low], nums[high] = nums[high], nums[low]
	}
	return nums
}
func canDivideBy3(x int) bool {
	return x%3 == 0
}
func isEven(x int) bool {
	return x%2 == 0
}

func main() {
	d := []int{1, 2, 3, 4, 5, 6, 7, 8, 9}
	fmt.Println(exchange(d))
	fmt.Println(exchangeByFunc(d, canDivideBy3))
	fmt.Println(exchangeByFunc(d, isEven))
}
