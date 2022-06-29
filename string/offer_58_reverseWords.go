package main

import (
	"fmt"
	"strings"
)

// strings.Split的缺点是，无法处理连续的分隔符
func reverseWordsTmp(s string) string {
	tmpStrs := strings.Split(s, " ")
	var res string
	for i := len(tmpStrs) - 1; i >= 0; i-- {
		res += tmpStrs[i] + " "
	}
	// 去掉末尾的空格
	res = res[:len(res)-1]

	return res
}

// 手动实现一遍分隔
func reverseWords(s string) string {
	var low, high int
	var res string
	length := len(s)
	// 特殊情况，输入空字符串
	if length == 0 {
		return s
	}

	for low < length {
		if s[low] == ' ' {
			low++
			continue
		} else {
			high = low
			// 所有移动指针的地方，一定要注意不能超出下标
			for high < length && s[high] != ' ' {
				high++
			}
			res = s[low:high] + " " + res
			low = high
		}
	}
	// 去掉末尾空格
	// 也要注意特殊情况，全为空格时，res去掉末尾空格会超出上标
	if len(res) > 0 {
		res = res[:len(res)-1]
	}

	return res
}

// 骚方法。strings.Fields方法可以切分连续空格，然后交换位置即可
func reverseWordsII(s string) string {
	strs := strings.Fields(s)
	for i := 0; i < len(strs)/2; i++ {
		strs[i], strs[len(strs)-i-1] = strs[len(strs)-i-1], strs[i]
	}

	return strings.Join(strs, " ")
}

func main() {
	d := "  hello world!  "
	fmt.Println(reverseWordsII(d))
}
