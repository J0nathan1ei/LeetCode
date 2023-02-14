package main

import "fmt"

// 这里比较容易碰壁的一点是：
// m,n要对应的是j和i,
// 因为m是切片的长度，对应的是二维切片dp[i][j]的j
// n对应了一维切片的个数，所以对应的是i
func uniquePaths(m int, n int) int {
	if m == 1 && n == 1 {
		return 1
	}

	dp := make([][]int, n)
	for num := 0; num < n; num++ {
		dp[num] = make([]int, m)
	}

	dp[0][0] = 1
	var i, j int
	for ; i < n; i++ {
		for j = 0; j < m; j++ {
			if i == 0 && j == 0 {
				continue
			}
			var num1, num2 int
			if i-1 >= 0 {
				num1 = dp[i-1][j]
			}
			if j-1 >= 0 {
				num2 = dp[i][j-1]
			}
			dp[i][j] = num1 + num2
		}
	}
	return dp[n-1][m-1]
}

// 初始化时，先把边界上的所有值都初始化
// 后面就不用判断边界条件了
func uniquePathsII(m int, n int) int {
	dp := make([][]int, m)
	for i := range dp {
		dp[i] = make([]int, n)
		dp[i][0] = 1
	}
	for j := 0; j < n; j++ {
		dp[0][j] = 1
	}
	for i := 1; i < m; i++ {
		for j := 1; j < n; j++ {
			dp[i][j] = dp[i-1][j] + dp[i][j-1]
		}
	}
	return dp[m-1][n-1]
}

func main() {
	n := uniquePaths(3, 3)
	fmt.Println(n)
}
