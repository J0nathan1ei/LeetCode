package main

import "fmt"

// 关键是如何递归地判断，让所有子树都去掉不包含1的节点
// 看了下答案
// 时间：20Min
func pruneTree(root *TreeNode) *TreeNode {
	if root == nil {
		return nil
	}
	var dfs func(root *TreeNode) *TreeNode
	dfs = func(root *TreeNode) *TreeNode {
		if root == nil {
			return nil
		}
		root.Left = dfs(root.Left)
		root.Right = dfs(root.Right)
		// 不包含1的节点，可以递归地转化为0-nil-nil型节点
		if root.Left == nil && root.Right == nil && root.Val == 0 {
			return nil
		}
		// 不是0-nil-nil型，就返回原有的节点
		return root
	}
	return dfs(root)
}

func main() {
	fmt.Println()
}
