package main

import "fmt"
/*
全排列这种问题的解决办法还是回溯，这里回溯时注意传入的对象是原数组nums和现有结果path
注意原数组恢复时，元素要提前保存一份(cur)
*/
func permute(nums []int) [][]int {
	length := len(nums)
	var res [][]int
	var path []int
	var backTrack func(nums, path []int)

	backTrack = func(nums, path []int){
		if len(path) == length{
			tmp := make([]int, length)
			copy(tmp, path)
			res = append(res, tmp)
		}

		for i:= 0;i < len(nums);i++{
			cur:=nums[i]
			path = append(path,nums[i])
			nums = append(nums[:i], nums[i+1:]...)
			backTrack(nums,path)
			nums = append(nums[:i],append([]int{cur},nums[i:]...)...)
			path = path[:len(path)-1]
		}
	}

	backTrack(nums, path)
	return res
}

func main(){
	d := []int{1,2,3}
	res := permute(d)
	fmt.Println(res)
}
