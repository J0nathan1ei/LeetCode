package main

import "fmt"

// 因为不让用除法，所以正常思维去乘所有的因子的话，时间复杂度是N^2
// 看了下答案
// 想法比较特别，是把数组拆分为左右两边的内容，进行计算的
// 时间复杂度N，空间复杂度N
// 时间：13min30s
func constructArr(a []int) []int {
	length := len(a)
	if length <= 1 {
		return []int{}
	}
	res := make([]int, length)
	pre := make([]int, length)
	post := make([]int, length)

	pre[0] = a[0]
	post[length-1] = a[length-1]

	for i := 1; i < length; i++ {
		pre[i] = pre[i-1] * a[i]
	}
	for i := length - 2; i >= 0; i-- {
		post[i] = post[i+1] * a[i]
	}

	res[0] = post[1]
	res[length-1] = pre[length-2]

	for i := 1; i < length-1; i++ {
		res[i] = pre[i-1] * post[i+1]
	}
	return res
}

func main() {
	d := []int{1, 2, 3, 4, 5}
	fmt.Println(constructArr(d))
}
