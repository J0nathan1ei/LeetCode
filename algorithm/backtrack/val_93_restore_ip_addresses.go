package main

import (
	"fmt"
	"strconv"
)

/*
典型的回溯搜索法，
关键点有以下几个
1.点分ip的path，直接用string不好表示，用[]string切片来表示，最后再组装
2.ip地址合法性的判断，注意不能以0开头的数字，如012
3.append一个ip元素时，要对下一步进行判断，在进行backTrack

*/
func restoreIpAddresses(s string) []string {
	var res []string
	length := len(s)
	if length < 4 || length > 12 {
		return res
	}
	var path []string
	var backTrack func(start int, path []string)
	backTrack = func(start int, path []string) {
		if start == length && len(path) == 4 {
			tmp := path[0] + "." + path[1] + "." + path[2] + "." + path[3]
			res = append(res, tmp)
		}
		for i := start; i < length; i++ {
			path = append(path, s[start:i+1])
			// 下面这一步，是考虑的关键，
			if i+1-start <= 3 && len(path) <= 4 && check(s, start, i+1) {
				backTrack(i+1, path)
			} else {
				return
			}
			path = path[:len(path)-1]
		}
	}

	backTrack(0, path)
	return res
}

func check(s string, start, end int) bool {
	if end-start > 1 && s[start] == '0' {
		return false
	}
	num, _ := strconv.Atoi(s[start:end])
	if num > 255 {
		return false
	}
	return true
}

func main() {
	d := "10102"

	fmt.Println(restoreIpAddresses(d))
}
