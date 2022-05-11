package main

import "fmt"
/*
基本思路是采用中心扩展法，从字符串头遍历到尾，需要注意遍历时的中心
时间复杂度n^2，空间复杂度为1
*/
func longestPalindrome(s string) string {
	length := len(s)
	if length <= 1 {
		return s
	}

	var res string
	var tmpLen, resLen int
	for i := 0; i < length*2-1; i++ {
		mid1, mid2 := i/2, (i+1)/2
		if mid1 == mid2 {
			tmpLen = 1
		} else if mid2 > mid1 && s[mid2] == s[mid1] {
			tmpLen = 2
		}
		for mid1 >= 0 && mid2 < length && s[mid1] == s[mid2] {
			tmpLen += 2
			mid1--
			mid2++
		}
		if tmpLen > resLen {
			resLen = tmpLen
			res = string(s[mid1+1 : mid2])
		}
	}
	return res
}

func max(x, y int) int {
	if x > y {
		return x
	}
	return y
}

func main() {
	d := "babad"
	fmt.Println(longestPalindrome(d))
}
