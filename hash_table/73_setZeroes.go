package main

import "fmt"

// 难点在于，如何让将要置0的元素，不影响到已有的置0元素

// 使用标记矩阵，就不会影响原有数组的内容了
func setZeroes(matrix [][]int) {
	lenRow := len(matrix)
	lenCol := len(matrix[0])
	var firstRowZero, firstColZero bool

	// 这里单从写法上来看，稍微有点反直觉
	for i := 0; i < lenRow; i++ {
		if matrix[i][0] == 0 {
			firstColZero = true
		}
	}

	for i := 0; i < lenCol; i++ {
		if matrix[0][i] == 0 {
			firstRowZero = true
		}
	}

	for i := 0; i < lenRow; i++ {
		for j := 0; j < lenCol; j++ {
			if matrix[i][j] == 0 {
				matrix[0][j] = 0
				matrix[i][0] = 0
			}
		}
	}

	for i := 1; i < lenRow; i++ {
		for j := 1; j < lenCol; j++ {
			if matrix[i][0] == 0 || matrix[0][j] == 0 {
				matrix[i][j] = 0
			}
		}
	}

	if firstColZero {
		for i := 0; i < lenRow; i++ {
			matrix[i][0] = 0
		}
	}
	if firstRowZero {
		for i := 0; i < lenCol; i++ {
			matrix[0][i] = 0
		}
	}
}
func main() {
	fmt.Println()
}
