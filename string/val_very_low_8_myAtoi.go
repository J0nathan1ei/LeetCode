package main

import (
	"fmt"
	"math"
)

// 这道题很傻逼，字符串转数字，一定要按照 “空格 + 正负号 + 数字 ”的模式去匹配，顺序换了就不行，正负号出现在0前面就可以，出现在0后面就不行
// 几个应该返回为0的例子：
// "00000-42a1234"    "words and 987"    "+-12"

// 返回正常数字的例子
// "+1"    "3.14159"    "  -0012a42"

func myAtoi(s string) int {
	var res int
	var isPositive bool = true
	var isPositiveChanged bool = false
	for i := 0; i < len(s); i++ {
		if s[i] == '-' && !isPositiveChanged {
			isPositive = false
			isPositiveChanged = true
			continue
		} else if s[i] == '+' && !isPositiveChanged {
			isPositiveChanged = true
			continue
		}
		if s[i] != ' ' {
			if s[i] < '0' || s[i] > '9' {
				return 0
			}
			if s[i] == '0' {
				continue
			}
			for j := i; j < len(s); j++ {
				if s[j] < '0' || s[j] > '9' {
					if s[j] != ' ' && s[j] != '.' {
						return 0
					}
					break
				}
				res = res*10 + int(s[j]-'0')
				if res >= math.MaxInt32+1 && !isPositive {
					return math.MinInt32
				}
				if res >= math.MaxInt32 && isPositive {
					return math.MaxInt32
				}
			}
			if !isPositive {
				res *= -1
			}
			return res
		}
	}
	return 0
}

func myAtoi2(s string) int {
	abs, sign, i, n := 0, 1, 0, len(s)
	//丢弃无用的前导空格
	for i < n && s[i] == ' ' {
		i++
	}
	//标记正负号
	if i < n {
		if s[i] == '-' {
			sign = -1
			i++
		} else if s[i] == '+' {
			sign = 1
			i++
		}
	}
	for i < n && s[i] >= '0' && s[i] <= '9' {
		abs = 10*abs + int(s[i]-'0')  //字节 byte '0' == 48
		if sign*abs < math.MinInt32 { //整数超过 32 位有符号整数范围
			return math.MinInt32
		} else if sign*abs > math.MaxInt32 {
			return math.MaxInt32
		}
		i++
	}
	return sign * abs
}

func main() {
	d := myAtoi2("+-1")
	fmt.Println(d)
}
