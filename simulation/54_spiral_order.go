package main

import "fmt"

//// 一开始没有用isWalked矩阵，出现了往回走的现象
//func spiralOrder(matrix [][]int) []int {
//	m, n := len(matrix), len(matrix[0])
//	left, right, up, down := 0, n-1, 1, m-1
//	var res []int
//	var i, j int
//	res = append(res, matrix[0][0])
//	for {
//		// 这个if其实是多余的，一开始想的是为了控制后面的right--，但是无法右移时right--也不会影响到终止条件
//		if j < right {
//			for j < right {
//				j++
//				res = append(res, matrix[i][j])
//			}
//			right--
//		}
//
//		if i < down {
//			for i < down {
//				i++
//				res = append(res, matrix[i][j])
//			}
//			down--
//		}
//
//		if j > left {
//			for j > left {
//				j--
//				res = append(res, matrix[i][j])
//			}
//			left++
//		}
//
//		if i > up {
//			for i > up {
//				i--
//				res = append(res, matrix[i][j])
//			}
//			up++
//		}
//
//		if left > right && up > down {
//			break
//		}
//	}
//	return res
//}

// 使用了isWalked矩阵的方法，防止往回走
func spiralOrder(matrix [][]int) []int {
	m, n := len(matrix), len(matrix[0])
	left, right, up, down := 0, n-1, 1, m-1
	var res []int
	var i, j int
	isWalked := make([][]bool, m)
	for index := 0; index < m; index++ {
		isWalked[index] = make([]bool, n)
	}
	res = append(res, matrix[0][0])
	isWalked[0][0] = true
	for {
		for j < right && !isWalked[i][j+1] {
			j++
			isWalked[i][j] = true
			res = append(res, matrix[i][j])
		}
		right--

		for i < down && !isWalked[i+1][j] {
			i++
			isWalked[i][j] = true
			res = append(res, matrix[i][j])
		}
		down--

		for j > left && !isWalked[i][j-1] {
			j--
			isWalked[i][j] = true
			res = append(res, matrix[i][j])
		}
		left++

		for i > up && !isWalked[i-1][j] {
			i--
			isWalked[i][j] = true
			res = append(res, matrix[i][j])
		}
		up++

		if left > right && up > down {
			break
		}
	}
	return res
}

func main() {
	d := [][]int{
		[]int{1, 2, 3, 4},
		[]int{5, 6, 7, 8},
		[]int{9, 10, 11, 12},
	}
	fmt.Println(spiralOrder(d))
}
