package main

import "fmt"

// 正常的方法：用两个指针遍历两个字符串
// 注意特殊情况的判断：s为空字符串则为true
// 还有就是s的指针，要移动到最后一个字符之后，而不是最后一个字符就行，否则长度为1的会判断错误
func isSubsequence(s string, t string) bool {
	lengthS := len(s)
	lengthT := len(t)
	if lengthS > lengthT {
		return false
	} else if lengthS == 0 {
		return true
	}
	var i, j int
	for i = 0; i < lengthT; i++ {
		if t[i] == s[j] {
			j++
		}
		if j == lengthS {
			return true
		}
	}
	return false
}
func main() {
	a := "b"
	b := "c"
	fmt.Println(isSubsequence(a, b))
}
