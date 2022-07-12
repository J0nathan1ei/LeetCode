package main

import (
	"fmt"
	"sort"
)

func kClosest(points [][]int, k int) [][]int {
	if k > len(points) {
		k = len(points)
	}
	sort.Slice(points, func(i, j int) bool {
		return points[i][0]*points[i][0]+points[i][1]*points[i][1] < points[j][0]*points[j][0]+points[j][1]*points[j][1]
	})
	return points[:k]
}

func main() {
	// go的sort.slice是可以对多维数组排序的
	//a1 := []int{1, 9}
	//a2 := []int{3, 4}
	//c1 := [][]int{a1, a2}
	//fmt.Println(c1)
	//sort.Slice(c1, func(i, j int) bool {
	//	return c1[i][0]*c1[i][0]+c1[i][1]*c1[i][1] < c1[j][0]*c1[j][0]+c1[j][1]*c1[j][1]
	//})
	d := [][]int{
		[]int{3, 3},
		[]int{5, -1},
		[]int{-2, 4},
	}
	fmt.Println(kClosest(d, 2))
}
