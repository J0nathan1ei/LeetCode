package main

import "fmt"

// 模拟栈操作，时间复杂度N，空间复杂度N
// 时间:35min
// 初始提交版

// 个人简化版
func validateStackSequences(pushed []int, popped []int) bool {
	// 排除特殊情况
	length := len(pushed)
	if length != len(popped) {
		return false
	}

	// 常规情况，模拟栈操作
	var stack []int
	var i, j int
	for i < length && j < length {
		top := popped[j]
		// 入栈操作
		for ; i < length; i++ {
			stack = append(stack, pushed[i])
			if pushed[i] == top {
				i++
				break
			}
		}
		// 出栈操作
		for ; j < length && len(stack) > 0; j++ {
			if stack[len(stack)-1] != popped[j] {
				break
			}
			stack = stack[:len(stack)-1]
		}
	}

	// 因为一定是先入栈，再出栈，所以只需判断出栈完了没有即可
	return j == length
}
func main() {
	d1 := []int{1, 2, 3, 4, 5}
	d2 := []int{4, 5, 3, 1, 2}
	fmt.Println(validateStackSequences(d1, d2))
}
