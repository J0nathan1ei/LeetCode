package main

import (
	"container/list"
	"fmt"
)

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

/*
典型的层序遍历，因为是z字形，所以额外涉及了队列的操作，
从左向右时，队列要pushback，
而从右向左时，队列要pushfront
*/
func zigzagLevelOrder(root *TreeNode) [][]int {
	var res [][]int
	var path []int
	if root == nil {
		return res
	}
	queue := list.New()
	queue.PushBack(root)
	var n int
	for queue.Len() > 0 {
		length := queue.Len()
		n++
		if n%2 == 1 { // 从左向右时
			for i := 0; i < length; i++ {
				cur := queue.Remove(queue.Front()).(*TreeNode)
				path = append(path, cur.Val)
				if cur.Left != nil {
					queue.PushBack(cur.Left)
				}
				if cur.Right != nil {
					queue.PushBack(cur.Right)
				}
			}
		} else { // 从右向左时
			for i := 0; i < length; i++ {
				cur := queue.Remove(queue.Back()).(*TreeNode)
				path = append(path, cur.Val)
				if cur.Right != nil {
					queue.PushFront(cur.Right)
				}
				if cur.Left != nil {
					queue.PushFront(cur.Left)
				}
			}
		}
		res = append(res, path)
		path = []int{}
	}
	return res
}

func main() {
	d := &TreeNode{3,
		&TreeNode{9, nil, nil}, &TreeNode{20,
			&TreeNode{15, nil, nil}, &TreeNode{7, nil, nil}}}
	fmt.Println(zigzagLevelOrder(d))
}
