package main

import "fmt"

// 将两个数组都拆分为根、左、右三部分
// 拆分中序数组时，要遍历吗
// 如何根据拆分结果来构造树

//

func buildTree(preorder []int, inorder []int) *TreeNode {
	if len(preorder) < 1 || len(inorder) < 1 {
		return nil
	}
	root := &TreeNode{preorder[0], nil, nil}
	var i int
	for ; i < len(inorder); i++ {
		if inorder[i] == root.Val {
			break
		}
	}
	root.Left = buildTree(preorder[1:i+1], inorder[:i])
	root.Right = buildTree(preorder[i+1:], inorder[i+1:])

	return root
}

func main() {
	fmt.Println()
}
