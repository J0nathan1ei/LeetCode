package main

import (
	"fmt"
	"sort"
)

func numFriendRequests(ages []int) int {
	var res int
	length := len(ages)

	sort.Ints(ages)
	for i := length - 1; i >= 0; i++ {
		var j int
		for j = 0; j < i; j++ {
			if ages[j] > ages[i]/2+7 || ages[j] >= ages[i] {
				break
			}
		}
		res += i - j
	}
	return res
}
func main() {
	d := []int{16, 16}
	fmt.Println(numFriendRequests(d))
}
