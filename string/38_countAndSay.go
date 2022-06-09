package main

import (
	"fmt"
	"strconv"
)

// 简单的字符串模拟
func countAndSay(n int) string {
	if n == 1 {
		return "1"
	}
	res := make([]string, n+1)
	res[1] = "1"
	for i := 2; i <= n; i++ {
		tmpRes := ""
		for j := 0; j < len(res[i-1]); j++ {
			cur := res[i-1][j]
			cnt := 0
			for j < len(res[i-1]) && res[i-1][j] == cur {
				cnt++
				j++
			}
			j--
			tmpRes += strconv.Itoa(cnt) + string(cur)
		}
		res[i] = tmpRes
	}
	return res[n]
}
func main() {
	fmt.Println(countAndSay(4))
}
