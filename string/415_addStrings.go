package main

import "fmt"

// 简单的字符串相加模拟
// 注意两个字符串的遍历，是在或条件中，一起遍历完的

func addStringsDiy(num1 string, num2 string) string {
	var res string

	m, n := len(num1)-1, len(num2)-1
	var cnt int
	for m >= 0 || n >= 0 {
		var n1, n2, sum int
		if m >= 0 {
			n1 = int(num1[m] - '0')
			m--
		}
		if n >= 0 {
			n2 = int(num2[n] - '0')
			n--
		}
		sum = n1 + n2 + cnt
		cnt = 0
		if sum >= 10 {
			cnt = 1
			sum -= 10
		}
		res = string('0'+sum) + res
	}
	if cnt == 1 {
		res = "1" + res
	}
	return res
}

func main() {
	a := "222333"
	b := "777662"
	fmt.Println(addStringsDiy(a, b))
}
