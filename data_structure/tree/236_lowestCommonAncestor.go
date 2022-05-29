package main

import "fmt"

// 首先是，找父节点时应该选取怎样的遍历方式？
// 其次，如何判断当前节点是两个节点的父节点？
// 如何开始树的递归

// 采用后序遍历，因为后序遍历就是从下到上的回溯
// 如果出现一个节点，它的左子树出现了P，而且右子树出现了Q
// 递归三部曲：
// (1)参数；   (2)终止条件；   (3)单层递归逻辑
//
func lowestCommonAncestor(root, p, q *TreeNode) *TreeNode {
	if root == nil || root == p || root == q {
		return root
	}

	left := lowestCommonAncestor(root.Left, p, q)
	right := lowestCommonAncestor(root.Right, p, q)

	if left != nil && right != nil {
		return root
	} else if left != nil {
		return left
	} else {
		return right
	}
}
func main() {
	fmt.Println()
}
