package main

import "fmt"

// 注意匹配时，有个比较特别的要求：
// 如果出现A为nil，B不为nil则不匹配
// 如果出现A不为nil，B为nil则为匹配
func isSubStructure(A *TreeNode, B *TreeNode) bool {
	if A == nil || B == nil {
		return false
	}
	var res bool = false
	var dfs func(root *TreeNode)
	dfs = func(root *TreeNode) {
		if root == nil {
			return
		}

		if root.Val == B.Val {
			res = match(root, B)
		}

		dfs(root.Left)
		dfs(root.Right)
	}
	dfs(A)
	return res
}

func match(A *TreeNode, B *TreeNode) bool {
	if A == nil && B == nil {
		return true
	} else if A == nil {
		return false
	} else if B == nil {
		return true
	}

	if A.Val != B.Val {
		return false
	}
	return match(A.Left, B.Left) && match(A.Right, B.Right)
}

func main() {
	fmt.Println()
}
