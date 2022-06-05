package main

import "fmt"

type BSTIterator struct {
	index       int
	inorderPath []int
}

func inorderInit(root *TreeNode) []int {
	var res []int
	var dfs func(root *TreeNode)
	dfs = func(root *TreeNode) {
		if root == nil {
			return
		}
		if root.Left != nil {
			dfs(root.Left)
		}
		res = append(res, root.Val)
		if root.Right != nil {
			dfs(root.Right)
		}
	}
	dfs(root)
	return res
}
func Constructor(root *TreeNode) BSTIterator {
	return BSTIterator{
		0,
		inorderInit(root),
	}
}

func (this *BSTIterator) Next() int {
	res := this.inorderPath[this.index]
	this.index++
	return res
}

func (this *BSTIterator) HasNext() bool {
	return this.index < len(this.inorderPath)
}

func main() {
	fmt.Println()
}
