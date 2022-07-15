package main

import "fmt"

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// 看了下题解
// 关键的地方是确立垂直遍历中，每个节点的顺序

func main() {
	d := &TreeNode{3,
		&TreeNode{9, nil, nil}, &TreeNode{20,
			&TreeNode{15, nil, nil}, &TreeNode{7, nil, nil}}}
	fmt.Println((d))
}
