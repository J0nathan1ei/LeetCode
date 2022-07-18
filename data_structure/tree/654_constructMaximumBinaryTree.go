package main

import "fmt"

//type TreeNode struct {
//	Val   int
//	Left  *TreeNode
//	Right *TreeNode
//}

// 典型的递归构建
// 时间：15min内搞定
func constructMaximumBinaryTree(nums []int) *TreeNode {
	if len(nums) == 0 {
		return nil
	}

	res := &TreeNode{0, nil, nil}
	index := maxIndex(nums)
	res.Val = nums[index]
	res.Left = constructMaximumBinaryTree(nums[:index])
	res.Right = constructMaximumBinaryTree(nums[index+1:])

	return res
}

func maxIndex(nums []int) int {
	var res int
	for i := 0; i < len(nums); i++ {
		if nums[i] > nums[res] {
			res = i
		}
	}
	return res
}
func main() {
	d := []int{3, 2, 1, 6, 0, 5}
	r := constructMaximumBinaryTree(d)
	fmt.Println(r.Val)
}
