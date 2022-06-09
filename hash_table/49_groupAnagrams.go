package main

import (
	"fmt"
)

// 基本思路是将单个字符串排序，
// 但是让go这么表示的话，有些费劲

// go可以把定长的数组作为key
// 这个特性注意下，估计是不让用不定长的元素作为key
func groupAnagrams(strs []string) [][]string {
	var res [][]string
	record := make(map[[26]int][]string)
	for _, val := range strs {
		tmpKey := [26]int{}
		for _, char := range val {
			tmpKey[char-'a']++
		}
		record[tmpKey] = append(record[tmpKey], val)
	}
	for _, val := range record {
		res = append(res, val)
	}
	return res
}

func main() {
	fmt.Println()
}
