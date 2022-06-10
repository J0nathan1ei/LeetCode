package main

import "fmt"

func jump(nums []int) int {
	length := len(nums)
	if length <= 1 {
		return 0
	}

	var res int
	for i := 0; i < len(nums); i++ {
		nums[i] += i
	}
	var low, high, tmpMax int
	for low = 0; low < length; low++ {
		res++
		for high = low; high <= nums[low] && high < len(nums); high++ {
			// 到达终点
			if nums[high] >= length-1 {
				// 每开始一趟循环都要加1，但是第一次的循环不用加1
				if high != 0 {
					res++
				}
				return res
			}
			tmpMax = max(tmpMax, high, nums)
		}
		// 将low移动到，可以产生最大位移的地方
		low = tmpMax - 1
	}
	return res
}

func max(x, y int, nums []int) int {
	if nums[x] > nums[y] {
		return x
	}
	return y
}
func main() {
	d := []int{2, 3, 1, 1, 4}
	fmt.Println(jump(d))
	d = []int{1, 2}
	fmt.Println(jump(d))
	d = []int{1, 1, 1, 1, 1}
	fmt.Println(jump(d))
	d = []int{1, 2, 1, 1, 1}
	fmt.Println(jump(d))

}
