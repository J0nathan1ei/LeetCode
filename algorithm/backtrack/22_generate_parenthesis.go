package main

import "fmt"

/*
典型的回溯枚举法，但是这种成对的操作第一次接触，需要对path两次改动
里面有个小点：回溯加右括号时要判断 right < left，表示左括号多的时候，才选择加右括号

*/
func generateParenthesis(n int) []string {
	var res []string
	var path string
	var backTrack func(left, right int, path string)
	backTrack = func(left, right int, path string) {
		if len(path) == 2*n {
			res = append(res, path)
			return
		}
		if left < n {
			path += "("
			backTrack(left+1, right, path)
			path = path[:len(path)-1]
		}
		if right < left {
			path += ")"
			backTrack(left, right+1, path)
			path = path[:len(path)-1]
		}
	}
	backTrack(0, 0, path)
	return res
}

func main() {
	fmt.Println(generateParenthesis(3))
}
