package main

import (
	"fmt"
	"sort"
)

// 难点在于如何对path进行去重
func combinationSum2(candidates []int, target int) [][]int {
	var res [][]int
	var path []int
	var backTrack func(start, sum int, path []int)
	sort.Ints(candidates)
	backTrack = func(start, sum int, path []int) {
		if sum == target {
			tmp := make([]int, len(path))
			copy(tmp, path)
			res = append(res, tmp)
			return
		} else if sum > target {
			return
		}

		for i := start; i < len(candidates); i++ {
			// i > start 这个条件很关键，代表当前第i个元素已经被放弃了，因此跳过重复的
			if i > start && candidates[i] == candidates[i-1] {
				continue
			}
			path = append(path, candidates[i])
			sum += candidates[i]
			backTrack(i+1, sum, path)
			sum -= candidates[i]
			path = path[:len(path)-1]

		}
	}
	backTrack(0, 0, path)
	return res
}
func main() {
	d := []int{1, 1, 1, 2}
	r := 3
	fmt.Println(combinationSum2(d, r))
}
