package main

import (
	"fmt"
	"math"
)

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func isValidBST(root *TreeNode) bool {
	var res bool = true
	var dfs func(root *TreeNode, low, high int)
	dfs = func(root *TreeNode, low, high int) {
		if root == nil {
			return
		}

		if root.Val <= low || root.Val >= high {
			res = false
			return
		}

		dfs(root.Left, low, root.Val)
		dfs(root.Right, root.Val, high)
	}

	dfs(root, math.MinInt32-1, math.MaxInt32+1)
	return res
}

func main() {
	fmt.Println(math.MinInt32)
}
