package main

import "fmt"

// 因为其他的都出现了两次
// 所以可以利用异或运算
func singleNumber(nums []int) int {
	var res int = nums[0]
	for i := 1; i < len(nums); i++ {
		res = res ^ nums[i]
	}
	return res
}

func main() {
	d := []int{1}
	fmt.Println(singleNumber(d))
}
