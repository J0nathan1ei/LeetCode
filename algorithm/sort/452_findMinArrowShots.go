package main

import (
	"fmt"
	"sort"
)

// 排序后取公共空间
//// points长度达到十万时，超时了
//func findMinArrowShots(points [][]int) int {
//	if len(points) == 1 {
//		return 1
//	}
//	sort.Slice(points, func(i, j int) bool {
//		return points[i][0] < points[j][0]
//	})
//	for i := 1; i < len(points); i++ {
//		if points[i-1][1] >= points[i][0] {
//			points[i][1] = min(points[i][1], points[i-1][1])
//			points = append(points[:i-1], points[i:]...)
//			i--
//		}
//	}
//	return len(points)
//}

func findMinArrowShots(points [][]int) int {
	if len(points) == 1 {
		return 1
	}
	sort.Slice(points, func(i, j int) bool {
		return points[i][0] < points[j][0]
	})
	res := 1
	maxRight := points[0][1]
	for i := 1; i < len(points); i++ {
		if maxRight >= points[i][0] {
			// 这里一开始写成了maxRight = min(points[i][1], points[i-1][1])，这样没有考虑到maxRight的更新保留
			// 卡了比较久
			maxRight = min(maxRight, points[i][1])
		} else {
			maxRight = points[i][1]
			res++
		}
	}
	return res
}
func min(x, y int) int {
	if x < y {
		return x
	}
	return y
}
func main() {
	d := [][]int{
		[]int{3, 9},
		[]int{7, 12},
		[]int{3, 8},
		[]int{6, 8},
		[]int{9, 10},
		[]int{2, 9},
		[]int{0, 9},
		[]int{3, 9},
		[]int{0, 6},
		[]int{2, 8},
	}

	fmt.Println(findMinArrowShots(d))
}
