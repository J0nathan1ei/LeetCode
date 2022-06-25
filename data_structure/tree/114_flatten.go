package main

import (
	"fmt"
)

//type TreeNode struct {
//	Val   int
//	Left  *TreeNode
//	Right *TreeNode
//}

func flattenII(root *TreeNode) {
	if root == nil {
		return
	}
	cur := root
	for cur != nil {
		if cur.Left != nil {
			node := getRight(cur.Left)
			node.Right = cur.Right
			cur.Left, cur.Right = nil, cur.Left
		}
		cur = cur.Right
	}
	return
}
func getRight(root *TreeNode) *TreeNode {
	if root.Left == nil && root.Right == nil {
		return root
	}
	if root.Right != nil {
		return getRight(root.Right)
	}
	if root.Left != nil {
		return getRight(root.Left)
	}
	return nil
}

func main() {
	d := &TreeNode{1,
		&TreeNode{2, &TreeNode{3, nil, nil}, &TreeNode{4, nil, nil}},
		&TreeNode{5, nil, &TreeNode{6, nil, nil}}}
	flattenII(d)
	fmt.Println(d.Val)
}
