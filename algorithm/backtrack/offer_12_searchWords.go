package main

import "fmt"

// 模拟走路的选择方向
var directions [4][2]int = [4][2]int{[2]int{-1, 0}, [2]int{1, 0}, [2]int{0, 1}, [2]int{0, -1}}

func exist(board [][]byte, word string) bool {
	var (
		visited                [][]bool
		backTrack              func(i, j int, index int) bool
		length, lenRow, lenCol int = len(word), len(board), len(board[0])
	)
	// 使用visited矩阵的目的是，防止重复走
	visited = make([][]bool, lenRow)
	for i := 0; i < lenRow; i++ {
		visited[i] = make([]bool, lenCol)
	}

	backTrack = func(i, j int, index int) bool {
		visited[i][j] = true
		// 当前函数退出前，要重新将visited置为false，使之只在自己的递归中生效
		defer func() { visited[i][j] = false }()

		if board[i][j] != word[index] {
			return false
		}
		// length-1表示到达了最终点
		if index == length-1 {
			return true
		}
		for _, val := range directions {
			if i+val[0] < 0 || i+val[0] >= lenRow ||
				j+val[1] < 0 || j+val[1] >= lenCol ||
				visited[i+val[0]][j+val[1]] {
				continue
			}
			if backTrack(i+val[0], j+val[1], index+1) {
				return true
			}
		}
		return false
	}

	// 回溯的范围是所有节点，但是可以提前减枝出来
	for i := 0; i < lenRow; i++ {
		for j := 0; j < lenCol; j++ {
			if backTrack(i, j, 0) {
				return true
			}
		}
	}
	return false
}

func main() {
	d := [][]byte{
		[]byte{'c', 'a', 'a'},
		[]byte{'a', 'a', 'a'},
		[]byte{'b', 'c', 'd'},
	}
	fmt.Println(exist(d, "aab"))
}
