package main

import "fmt"

// 一开始7min没想出来，看了下答案
// 它的基本思路是，首先将问题转化为和为0的问题，同时用哈希表记录好遍历过程中的这些和
// 如果出现了相同的和，说明中间部分肯定和为0，即0和1的个数相同
func findMaxLengthII(nums []int) (maxLength int) {
	mp := map[int]int{0: -1}
	counter := 0
	for i, num := range nums {
		if num == 1 {
			counter++
		} else {
			counter--
		}
		if prevIndex, has := mp[counter]; has {
			maxLength = max(maxLength, i-prevIndex)
		} else {
			mp[counter] = i
		}
	}
	return
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func main() {
	d := []int{0, 1, 0}
	fmt.Println(findMaxLengthII(d))
}
