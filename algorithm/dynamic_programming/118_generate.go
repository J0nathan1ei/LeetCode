package main

import "fmt"

func generate(numRows int) [][]int {
	var res [][]int
	for i := 1; i <= numRows; i++ {
		var path []int
		for j := 0; j < i; j++ {
			if j == 0 || j == i-1 {
				path = append(path, 1)
			} else {
				path = append(path, res[i-2][j-1]+res[i-2][j])
			}
		}
		res = append(res, path)
	}
	return res
}

func main() {
	fmt.Println(generate(5))
}
