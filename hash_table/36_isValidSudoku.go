package main

import "fmt"

// 模拟题，利用了哈希表
// 难点在于如何想到用哈希表，来模拟这个3x3x9的数独
func isValidSudoku(board [][]byte) bool {
	var rows, cols [9][9]int
	var subArray [3][3][9]int
	for i, row := range board {
		for j, col := range row {
			if col == '.' {
				continue
			}
			position := col - '1'
			rows[i][position]++
			cols[j][position]++
			subArray[i/3][j/3][position]++
			if rows[i][position] > 1 || cols[j][position] > 1 || subArray[i/3][j/3][position] > 1 {
				return false
			}
		}
	}
	return true
}

func main() {
	fmt.Println()
}
