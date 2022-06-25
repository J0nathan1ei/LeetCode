package main

import "fmt"

func minPathSum(grid [][]int) int {
	var (
		lenRow int = len(grid)
		lenCol int = len(grid[0])
	)

	for i := 0; i < lenRow; i++ {
		for j := 0; j < lenCol; j++ {
			var up, left int = -1, -1
			if i-1 >= 0 {
				up = grid[i-1][j]
			}
			if j-1 >= 0 {
				left = grid[i][j-1]
			}
			grid[i][j] = grid[i][j] + minSub(up, left)
		}
	}

	return grid[lenRow-1][lenCol-1]
}

func minSub(x, y int) int {
	if x == -1 && y == -1 {
		return 0
	} else if x == -1 {
		return y
	} else if y == -1 {
		return x
	}
	if x < y {
		return x
	}
	return y
}
func main() {
	d := [][]int{[]int{1, 2, 3}, []int{4, 5, 6}}
	fmt.Println(minPathSum(d))
}
