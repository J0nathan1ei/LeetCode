package main

import "fmt"

func moveZeroes(nums []int) {
	var length int = len(nums)
	if length == 0 {
		return
	}

	var i, j int
	// i指向第一个0
	for ; i < length; i++ {
		if nums[i] == 0 {
			break
		}
	}
	for j = i; j < length; j++ {
		if nums[j] == 0 {
			continue
		}
		nums[i] = nums[j]
		i++
	}
	for ; i < length; i++ {
		nums[i] = 0
	}
	return
}

// 学习下题解方法：
// 代码更加简洁，他只从快指针的角度，写了一个循环
// 而且因为是数值交换，就没有了最后的批量赋0操作
// 左指针跟随右指针移动，右指针碰见0，则左指针就不动了
// 主体思路就是：右指针遍历一遍找到所有符合条件的非零数，跟0交换
// 左指针的处理也很关键，
func moveZeroesII(nums []int) {
	left, right, n := 0, 0, len(nums)
	for right < n {
		if nums[right] != 0 {
			nums[left], nums[right] = nums[right], nums[left]
			left++
		}
		right++
	}
}

func main() {
	d := []int{0, 0, 0, 0, 0}
	moveZeroes(d)
	fmt.Println(d)
}
