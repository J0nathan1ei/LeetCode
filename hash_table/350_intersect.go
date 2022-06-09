package main

import "fmt"

func intersect(nums1 []int, nums2 []int) []int {
	var res []int
	m := make(map[int]int)
	for _, val := range nums1 {
		if _, ok := m[val]; !ok {
			m[val] = 1
		} else {
			m[val]++
		}
	}

	for _, val := range nums2 {
		if m[val] > 0 {
			res = append(res, val)
			m[val]--
		}
	}
	return res
}
func main() {
	fmt.Println()
}
