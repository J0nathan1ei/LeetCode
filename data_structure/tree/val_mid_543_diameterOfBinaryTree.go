package main

import "fmt"

// 注意这个终止条件的利用，返回0值，但是在外部会返回+1

func diameterOfBinaryTree(root *TreeNode) int {
	if root == nil {
		return 0
	}
	var res int
	var dfs func(root *TreeNode) int
	dfs = func(root *TreeNode) int {
		if root == nil {
			return 0
		}
		left := dfs(root.Left)
		right := dfs(root.Right)
		res = max(res, left+right)
		return 1 + max(left, right)
	}

	dfs(root)
	return res
}

func max(x, y int) int {
	if x > y {
		return x
	}
	return y
}

func main() {
	fmt.Println()
}
