package main

import "fmt"

// 本来以为是回溯题，想了一半，发现是数学题
// 想到数学题的时候，有去重的情况，但是题目里面其实不去重，[7,7,7,7]子序列是3个
// 时间：30min
// 使用path数组+数学解法的空间复杂度是N
func numberOfArithmeticSlices(nums []int) int {
	length := len(nums)
	if length < 3 {
		return 0
	}
	var res int
	var path []int
	path = append(path, nums[0], nums[1])

	for i := 2; i < length; i++ {
		// 发现当前元素满足等差，则加入path
		if nums[i]-path[len(path)-1] == path[1]-path[0] {
			path = append(path, nums[i])
		} else {
			// 不满足等差时，如果成员数为2，则重新开始下一轮
			if len(path) == 2 {
				path = []int{path[1], nums[i]}
				continue
			}
			// 成员数大于3时， 统计一下已有的子序列数（数学方法）
			res += (len(path) - 1) * (len(path) - 2) / 2
			path = []int{path[len(path)-1], nums[i]}
		}
		// 到达结尾, 且path成员大于3
		if i == length-1 && len(path) >= 3 {
			res += (len(path) - 1) * (len(path) - 2) / 2
			path = []int{path[len(path)-1], nums[i]}
		}
	}
	return res
}

// 官方的解法跟我的想法大致一样，也是数学解法，但是没有用到path数组，所以空间复杂度是1
// 不过实际内存就比我想的小了100KB
func numberOfArithmeticSlicesII(nums []int) (ans int) {
	n := len(nums)
	if n == 1 {
		return
	}

	d, t := nums[0]-nums[1], 0
	// 因为等差数列的长度至少为 3，所以可以从 i=2 开始枚举
	for i := 2; i < n; i++ {
		if nums[i-1]-nums[i] == d {
			t++
		} else {
			d, t = nums[i-1]-nums[i], 0
		}
		ans += t
	}
	return
}

func main() {
	d := []int{1, 2, 3}
	fmt.Println(numberOfArithmeticSlices(d))
}
