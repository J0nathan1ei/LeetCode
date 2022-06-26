package main

import (
	"fmt"
	"sort"
)

// 标准的回溯方法，注意去重的方法
// 13min解决
func subsetsWithDup(nums []int) [][]int {
	var (
		path      []int
		res       [][]int
		length    = len(nums)
		backTrack func(start int, path []int)
	)
	sort.Ints(nums)
	backTrack = func(start int, path []int) {
		tmp := make([]int, len(path))
		copy(tmp, path)
		res = append(res, tmp)

		for i := start; i < length; i++ {
			if i > start && nums[i] == nums[i-1] {
				continue
			}
			path = append(path, nums[i])
			backTrack(i+1, path)
			path = path[:len(path)-1]
		}
	}
	backTrack(0, path)

	return res
}
func main() {
	d := []int{0}
	fmt.Println(subsetsWithDup(d))
}
