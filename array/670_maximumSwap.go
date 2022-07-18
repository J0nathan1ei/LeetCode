package main

import "fmt"

// 基本的思路，就是把左边较小的，和右边最大的交换
// 想到一种N^2的方法，左右两个指针向中间移动
// 带调试25min
func maximumSwap(num int) int {
	var nums []int
	var res int
	for num != 0 {
		tmp := num % 10
		num /= 10
		nums = append([]int{tmp}, nums...)
	}
	low, high := 0, len(nums)-1
	for low < high {
		tmpHigh := high
		maxIndex := high
		for low < tmpHigh {
			tmpHigh--
			maxIndex = maxNumsIndex(maxIndex, tmpHigh, nums)
		}
		if nums[maxIndex] > nums[low] {
			// 最左边的low，只要发现满足交换的数字，就交换并退出
			nums[maxIndex], nums[low] = nums[low], nums[maxIndex]
			break
		}
		low++
	}
	// 还原交换后的结果
	for i := 0; i < len(nums); i++ {
		res = res*10 + nums[i]
	}
	return res
}

func maxNumsIndex(x, y int, nums []int) int {
	// 这里要添加一个等号的原因是，如果右边都是大于low，但是相等的话，优先选择更右边的数
	// 因为位数更低，换位后对原数字影响更小
	if nums[x] >= nums[y] {
		return x
	}
	return y
}

func main() {
	d := 1993
	fmt.Println(maximumSwap(d))
}
