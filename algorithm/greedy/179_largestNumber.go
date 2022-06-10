package main

import (
	"fmt"
	"sort"
	"strconv"
)

//// 一开始的想法，按照字符串转成数字的各种可能性排序
//func largestNumber(nums []int) string {
//	sort.Slice(nums, func(i, j int) bool {
//		return maxNumStr(nums[i], nums[j])
//	})
//	var res string
//	for i := 0; i < len(nums); i++ {
//		res += strconv.Itoa(nums[i])
//	}
//	return res
//}
//func maxNumStr(x, y int) bool {
//	strX, strY := strconv.Itoa(x), strconv.Itoa(y)
//	for i := 0; i < len(strX) && i < len(strY); i++ {
//		if strX[i] > strY[i] {
//			return true
//		} else if strX[i] < strY[i] {
//			return false
//		}
//
//		if i == len(strX)-1 {
//			if len(strY) > len(strX) && strY[i+1] > strX[i] {
//				return false
//			}
//			return true
//		} else if i == len(strY)-1 {
//			if len(strX) > len(strY) && strX[i+1] > strY[i] {
//				return true
//			}
//			return false
//		}
//	}
//	return false
//}

// 比较牛批的做法，直接字符串排序
func largestNumber(nums []int) string {
	var res string
	length := len(nums)
	path := make([]string, length)
	for i := 0; i < length; i++ {
		path[i] = strconv.Itoa(nums[i])
	}
	sort.Slice(path, func(i, j int) bool {
		return path[i]+path[j] > path[j]+path[i]
	})
	for i := 0; i < length; i++ {
		res += path[i]
	}
	if res[0] == '0' {
		return "0"
	}
	return res
}
func main() {
	d := []int{3, 30, 34, 5, 9}
	fmt.Println(largestNumber(d))
}
