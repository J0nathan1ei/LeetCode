package main

import "fmt"

// 逐个移动的话，时间复杂度有些大，一次移动好之后赋值到矩阵中
// 空间复杂度N，时间复杂度N
// Time:12min
func shiftGrid(grid [][]int, k int) [][]int {
	lenCol, lenRow := len(grid), len(grid[0])
	total := lenCol * lenRow
	k = k % total
	if k == 0 {
		return grid
	}
	var nums []int
	for i := 0; i < lenCol; i++ {
		nums = append(nums, grid[i]...)
	}

	// 一维数组内移位，移动完之后赋值
	nums = append(nums[total-k:], nums[:total-k]...)
	for i := 0; i < lenCol; i++ {
		grid[i] = nums[i*lenRow : (i+1)*lenRow]
	}
	return grid
}
func main() {
	d := [][]int{
		[]int{1, 2, 3},
		[]int{4, 5, 6},
		[]int{7, 8, 9},
	}
	fmt.Println(shiftGrid(d, 1))
}
