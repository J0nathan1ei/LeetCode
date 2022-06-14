package main

import "fmt"

// 初步判断回溯法
// 第一次提交，回溯的时候，漏判了一个条件
// 第二次提交回溯，对于"111111111111111111111111111111111111111111111"还是超时了
func numDecodings(s string) int {
	var (
		res, length int
		path        string
		backTrack   func(start int, path string)
	)
	length = len(s)
	backTrack = func(start int, path string) {
		if len(path) == length {
			res++
			return
		}
		for i := start; i < length; i++ {
			// 一开始回溯的时候，这里的条件忽略了只要前一个是1，后一个可以是任何数字
			if i < length-1 && ((s[i] == '1') || (s[i] == '2' && s[i+1] <= '6')) {
				path += string(s[i : i+2])
				backTrack(i+2, path)
				path = path[:len(path)-2]
			}
			if s[i] == '0' {
				continue
			}
			// 正常流程
			path += string(s[i])
			backTrack(i+1, path)
			path = path[:len(path)-1]
		}
	}
	backTrack(0, path)
	return res
}

// 官方动态规划法
// 怎么想到动态规划法的？
//
func numDecodingsII(s string) int {
	n := len(s)
	f := make([]int, n+1)
	f[0] = 1
	for i := 1; i <= n; i++ {
		if s[i-1] != '0' {
			f[i] += f[i-1]
		}
		if i > 1 && s[i-2] != '0' && ((s[i-2]-'0')*10+(s[i-1]-'0') <= 26) {
			f[i] += f[i-2]
		}
	}
	return f[n]
}

func main() {
	d := "2611055971756562"
	fmt.Println(numDecodingsII(d))
}
