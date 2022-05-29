package main

import "fmt"

//type TreeNode struct {
//	Val   int
//	Left  *TreeNode
//	Right *TreeNode
//}

// 如何进行路径搜索
// 明显是回溯，但是是树里的回溯
// copy函数的使用

func pathSum(root *TreeNode, targetSum int) [][]int {
	var res [][]int
	var path []int

	if root == nil {
		return res
	}

	var dfs func(sum int, root *TreeNode)
	dfs = func(sum int, root *TreeNode) {
		if root == nil {
			return
		}

		sum -= root.Val
		path = append(path, root.Val)
		if sum == 0 && root.Left == nil && root.Right == nil {
			tmp := make([]int, len(path))
			copy(tmp, path)
			res = append(res, tmp)
		}

		dfs(sum, root.Left)
		dfs(sum, root.Right)
		path = path[:len(path)-1]
	}

	dfs(targetSum, root)
	return res
}

func main() {
	d := []int{1, 2, 3}
	tmp := make([]int, len(d))
	copy(tmp, d)
	fmt.Println(tmp)
}
