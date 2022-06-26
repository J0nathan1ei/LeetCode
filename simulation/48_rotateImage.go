package main

import "fmt"

// 虽然是模拟，但是主要考察的是，在模拟旋转过程中的细节处理：
// 旋转过程中，不能把最后一个元素也旋转了，否则相当于把第一个元素旋转了两遍
func rotate(matrix [][]int) {
	length := len(matrix)
	if length <= 1 {
		return
	}
	for i := 0; i < length/2; i++ {
		for j := i; j < length-i-1; j++ {
			matrix[i][j], matrix[j][length-1-i], matrix[length-1-i][length-1-j], matrix[length-1-j][i] =
				matrix[length-1-j][i], matrix[i][j], matrix[j][length-1-i], matrix[length-1-i][length-1-j]
		}
	}
	return
}

func main() {
	d := [][]int{
		[]int{5, 1, 9, 11},
		[]int{2, 4, 8, 10},
		[]int{13, 3, 6, 7},
		[]int{15, 14, 12, 16},
	}
	rotate(d)
	fmt.Println(d)
}
