package main

import "fmt"

// 想到两个数组互相移动，将较大的数组后面拼接一部分，来让较小的数组往后移动
// 时间复杂度(M+N)*min(M,N),空间复杂度1
// 时间 50min+
// 有几处细节仍然需要注意：
// 在长数组后面追加内容的时候，要防止结果首尾相连的情况
// 遍历长数组的时候，不能超过长数组的长度
func findLength(nums1 []int, nums2 []int) int {
	var res, tmpRes int
	var minNums, maxNums []int
	if len(nums1) < len(nums2) {
		minNums = nums1
		maxNums = nums2
	} else {
		minNums = nums2
		maxNums = nums1
	}
	maxNums = append(maxNums, maxNums[:len(minNums)]...)
	for i := 0; i < len(maxNums); i++ {
		tmpStep := 0
		tmpI := i
		// tmpI遍历长数组时，不能超过长数组的长度
		for j := 0; j < len(minNums) && tmpI < len(maxNums); j++ {
			if minNums[j] == maxNums[tmpI] {
				tmpStep++
				// 防止数组的首尾相连
				if tmpI == len(maxNums)-len(minNums) {
					tmpStep = 1
				}
			} else {
				tmpStep = 0
			}
			tmpI++
			tmpRes = max(tmpRes, tmpStep)
		}
		res = max(res, tmpRes)
	}
	return res
}
func min(x, y int) int {
	if x < y {
		return x
	}
	return y
}
func max(x, y int) int {
	if x > y {
		return x
	}
	return y
}

func main() {
	nums1 := []int{1, 0, 0, 0, 0}
	nums2 := []int{0, 0, 0, 0, 1}
	fmt.Println(findLength(nums1, nums2))
}
