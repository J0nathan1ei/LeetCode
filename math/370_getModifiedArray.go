package main

import "fmt"

// 差分数组，适合大量数据，要做相同的加减改动的情况
// 本质是以空间换时间
// 可参考 https://blog.csdn.net/qq_31601743/article/details/105352885
func getModifiedArray(length int, updates [][]int) []int {
	res := make([]int, length)
	for _, nums := range updates {
		res[nums[0]] += nums[2]
		if nums[1]+1 < length {
			res[nums[1]+1] -= nums[2]
		}
	}
	for i := 1; i < length; i++ {
		res[i] += res[i-1]
	}
	return res
}
func main() {
	d := [][]int{
		[]int{2, 4, 6},
		[]int{5, 6, 8},
		[]int{1, 9, -4},
	}
	fmt.Println(getModifiedArray(10, d))
}
