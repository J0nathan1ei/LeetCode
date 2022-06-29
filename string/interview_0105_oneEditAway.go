package main

import "fmt"

// 一次过
// 用时：刚好15min
func oneEditAway(first string, second string) bool {
	// 特殊情况，字符串相等
	if first == second {
		return true
	}
	// 特殊情况，字符串长度差超过了2
	len1, len2 := len(first), len(second)
	if abs(len1, len2) > 1 {
		return false
	}

	// 常规情况
	// 在遇到第一个不相同的字母之前，一直移动
	var i, j int
	for i < len1 && j < len2 {
		if first[i] != second[j] {
			break
		}
		i++
		j++
	}
	// 遇到第一个不同，跳到后面的字符
	if len1 == len2 {
		i++
		j++
	} else if len1 > len2 {
		i++
	} else if len2 > len1 {
		j++
	}

	// 如果后面再出现不同，就返回false
	for i < len1 && j < len2 {
		if first[i] != second[j] {
			return false
		}
		i++
		j++
	}
	return true
}

func abs(x, y int) int {
	if x-y > 0 {
		return x - y
	}
	return y - x
}
func main() {
	a := "pale"
	b := "ple"
	fmt.Println(oneEditAway(a, b))
}
