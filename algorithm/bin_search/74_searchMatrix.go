package main

import "fmt"

func searchMatrix(matrix [][]int, target int) bool {
	m, n := len(matrix), len(matrix[0])
	// 极端值判断
	if target < matrix[0][0] || target > matrix[m-1][n-1] {
		return false
	}
	// 开始内部搜索
	var row int
	for row = 0; row <= m-1; row++ {
		if matrix[row][0] == target {
			return true
		}
		if matrix[row][0] > target {
			break
		}
	}
	row--

	// 二分搜索
	var low, high = 0, n - 1
	for low <= high {
		mid := low + (high-low)/2
		if matrix[row][mid] == target {
			return true
		} else if matrix[row][mid] < target {
			low = mid + 1
		} else if matrix[row][mid] > target {
			high = mid - 1
		}
	}
	return false
}
func main() {
	d := [][]int{
		[]int{1, 3, 5, 7},
		[]int{10, 11, 16, 20},
		[]int{23, 30, 34, 60},
	}
	fmt.Println(searchMatrix(d, 60))
}
