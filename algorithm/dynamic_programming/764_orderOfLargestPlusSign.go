package main

import "fmt"

// 一开始没想到什么好的办法，只有暴力破解，遍历每一个元素，这样的话时间复杂度会达到N^3
// 看了下参考答案
// 时间 30min+
func orderOfLargestPlusSign(n int, mines [][]int) int {
	dp := make([][]int, n)
	for i := 0; i < n; i++ {
		dp[i] = make([]int, n)
	}
	// 标记一下dp中的mine
	for i := 0; i < len(mines); i++ {
		x, y := mines[i][0], mines[i][1]
		dp[x][y] = -1
	}

	var res, count int
	for i := 0; i < n; i++ {
		// 从左到右
		count = 0
		for j := 0; j < n; j++ {
			if dp[i][j] == -1 {
				count = 0
			} else {
				count++
				dp[i][j] = count
			}
		}

		// 从右到左
		count = 0
		for j := n - 1; j >= 0; j-- {
			if dp[i][j] == -1 {
				count = 0
			} else {
				count++
				dp[i][j] = min(dp[i][j], count)
			}
		}
	}

	for j := 0; j < n; j++ {
		//从上到下
		count = 0
		for i := 0; i < n; i++ {
			if dp[i][j] == -1 {
				count = 0
			} else {
				count++
				dp[i][j] = min(dp[i][j], count)
			}
		}

		//从下到上
		count = 0
		for i := n - 1; i >= 0; i-- {
			if dp[i][j] == -1 {
				count = 0
			} else {
				count++
				dp[i][j] = min(dp[i][j], count)
			}
			// 最后一次遍历，获取到答案
			res = max(res, dp[i][j])
		}
	}
	return res
}

//func min(x, y int) int {
//	if x < y {
//		return x
//	}
//	return y
//}
//func max(x, y int) int {
//	if x > y {
//		return x
//	}
//	return y
//}

func main() {
	n := 1
	mines := [][]int{
		[]int{0, 0},
	}
	fmt.Println(orderOfLargestPlusSign(n, mines))
}
