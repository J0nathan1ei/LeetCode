package main

import "fmt"

// 这种就算是算法思想了，C++的next_permutation字典序也是基于这个原理去做的
// 从后面寻找较小数、较大数并交换，然后将交换后的尾部序列重新reverse，就可得到下一个字典序
func nextPermutation(nums []int) {
	length := len(nums)
	if length == 1 {
		return
	}
	var i, j int
	for i = length - 2; i >= 0; i-- {
		if nums[i] < nums[i+1] {
			break
		}
	}
	if i >= 0 {
		j = length - 1
		for nums[j] <= nums[i] {
			j--
		}
		nums[i], nums[j] = nums[j], nums[i]
	}

	reverse(nums[i+1:])
}

func reverse(nums []int) {
	length := len(nums)
	for low, high := 0, length-1; low < high; {
		nums[low], nums[high] = nums[high], nums[low]
		low++
		high--
	}
}

func main() {
	d := []int{1, 3, 2}
	nextPermutation(d)
	fmt.Println(d)
}
