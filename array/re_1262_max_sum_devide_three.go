package main

import "fmt"

// 解法是有限状态机的思想，最大子和为3，它的前一个状态可以从余0、余1、余2的状态中推出，
// 因此计算出前一步的余0、余1、余2子和各自能够推导到哪一步的状态即可
func maxSumDivThree(nums []int) int {
	rest := [3]int{}
	for _, num := range nums {
		a := rest[0] + num
		b := rest[1] + num
		c := rest[2] + num
		rest[a%3] = max(rest[a%3], a)
		rest[b%3] = max(rest[b%3], b)
		rest[c%3] = max(rest[c%3], c)
	}
	return rest[0]
}

//func max(x, y int) int {
//	if x > y {
//		return x
//	}
//	return y
//}

func main() {
	d := []int{1, 2, 4}
	fmt.Println(maxSumDivThree(d))
}
