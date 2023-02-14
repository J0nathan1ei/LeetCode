package main

import "fmt"

func uniquePathsWithObstacles(obstacleGrid [][]int) int {
	n, m := len(obstacleGrid), len(obstacleGrid[0])
	var hasObstacle bool = false
	dp := make([][]int, n)
	for num := 0; num < n; num++ {
		dp[num] = make([]int, m)
		if obstacleGrid[num][0] == 1 {
			hasObstacle = true
		}
		if !hasObstacle {
			dp[num][0] = 1
		}
	}
	hasObstacle = false
	for num := 0; num < m; num++ {
		if obstacleGrid[0][num] == 1 {
			hasObstacle = true
		}
		if !hasObstacle {
			dp[0][num] = 1
		}
	}

	for i := 1; i < n; i++ {
		for j := 1; j < m; j++ {
			if obstacleGrid[i][j] == 1 {
				dp[i][j] = 0
			} else {
				dp[i][j] = dp[i-1][j] + dp[i][j-1]
			}
		}
	}
	return dp[n-1][m-1]
}

func main() {
	a := [][]int{
		[]int{0, 1},
		[]int{0, 0},
	}
	n := uniquePathsWithObstacles(a)
	fmt.Println(n)
}
