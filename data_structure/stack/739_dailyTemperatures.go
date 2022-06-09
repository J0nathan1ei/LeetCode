package main

import "fmt"

// 单调栈做法，比较巧妙
func dailyTemperatures(temperatures []int) []int {
	length := len(temperatures)
	res := make([]int, length)
	var stack []int

	for i := 0; i < length; i++ {
		temperature := temperatures[i]
		// 退栈: 当栈中有元素，且新元素不满足单调性，则进行退栈操作
		// 退栈时，将结果中的对应天数修改
		for len(stack) > 0 && temperature > temperatures[stack[len(stack)-1]] {
			prevIndex := len(stack) - 1
			res[stack[prevIndex]] = i - stack[prevIndex]
			stack = stack[:prevIndex]
		}
		// 退栈完毕，保留本次元素
		stack = append(stack, i)
	}
	return res
}

func main() {
	d := []int{73, 74, 75, 71, 69, 72, 76, 73}
	fmt.Println(dailyTemperatures(d))
}
