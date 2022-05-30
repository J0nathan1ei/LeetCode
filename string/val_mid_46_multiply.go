package main

import (
	"fmt"
	"math"
	"strconv"
)

// 不能采用string和int互转的方式，因为会超出int64的界限

// 应该使用字符串模拟乘法的形式来解决————乘一位+字符串加法

// 字符串模拟操作
func multiply(num1 string, num2 string) string {
	if num1 == "0" || num2 == "0" {
		return "0"
	}
	ans := "0"
	m, n := len(num1), len(num2)
	for i := n - 1; i >= 0; i-- {
		curr := ""
		add := 0
		for j := n - 1; j > i; j-- {
			curr += "0"
		}
		y := int(num2[i] - '0')
		for j := m - 1; j >= 0; j-- {
			x := int(num1[j] - '0')
			product := x*y + add
			curr = strconv.Itoa(product%10) + curr
			add = product / 10
		}
		for ; add != 0; add /= 10 {
			curr = strconv.Itoa(add%10) + curr
		}
		ans = addStrings(ans, curr)
	}
	return ans
}

func addStrings(num1, num2 string) string {
	i, j := len(num1)-1, len(num2)-1
	add := 0
	ans := ""
	for ; i >= 0 || j >= 0 || add != 0; i, j = i-1, j-1 {
		x, y := 0, 0
		if i >= 0 {
			x = int(num1[i] - '0')
		}
		if j >= 0 {
			y = int(num2[j] - '0')
		}
		result := x + y + add
		ans = strconv.Itoa(result%10) + ans
		add = result / 10
	}
	return ans
}

// 字符串Int互转
func multiply2(num1 string, num2 string) string {
	n1, n2 := stringToNum(num1), stringToNum(num2)
	return numToString(n1 * n2)
}

func stringToNum(s string) int {
	var res int
	for i := 0; i < len(s); i++ {
		res = res*10 + int(s[i]-'0')
	}
	return res
}
func numToString(num int) string {
	var res string
	for num > 0 {
		n := num % 10
		num /= 10
		res = string('0'+n) + res
	}
	if res == "" {
		return "0"
	}
	return res
}

func main() {
	d := multiply("5", "9223372036854775807")
	fmt.Println(d)
	fmt.Println(math.MaxInt64)
}
