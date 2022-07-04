package main

import "fmt"

// 想到了525的哈希表做法，但是没想清楚哈希表中应该存什么，
// 初步想法还是存和，打算用相减的方法来找出所需数组，但是问题是中间的数组表示出来，可能需要n^2的时间复杂度
// 原理同525，只是这次哈希表中记录的是余数，取余的作用在于去重，以使得哈希表命中
func checkSubarraySum(nums []int, k int) bool {
	return false
}

func main() {
	fmt.Println()
}
