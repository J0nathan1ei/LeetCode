package main

import "fmt"

// 二维地枚举所有的背包重量与物品摆放的关系
// 纵坐标是物品情况，横坐标是背包重量情况
func bagProblem(weight, value []int, bagWeight int) int {
	if len(weight) != len(value) {
		return -1
	}
	n := len(value)
	dp := make([][]int, n)
	for i := 0; i < n; i++ {
		dp[i] = make([]int, bagWeight+1)
	}
	for i := weight[0]; i < bagWeight; i++ {
		dp[0][i] = value[0]
	}

	for i := 1; i < n; i++ {
		for j := 0; j <= bagWeight; j++ {
			if weight[i] > j {
				dp[i][j] = dp[i-1][j]
			} else {
				dp[i][j] = max(dp[i-1][j], dp[i-1][j-weight[i]]+value[i])
			}
		}
	}

	return dp[n-1][bagWeight]
}

//func max(x, y int) int {
//	if x > y {
//		return x
//	}
//	return y
//}

func main() {
	weight := []int{1, 3, 4}
	value := []int{15, 20, 30}
	bagweight := 4
	res := bagProblem(weight, value, bagweight)
	fmt.Println(res)
}
