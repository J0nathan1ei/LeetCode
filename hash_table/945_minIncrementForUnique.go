package main

import (
	"fmt"
	"sort"
)

// 第一次想到的是哈希表
// 时间15min
// 但是遇到大数组的时候超时了
func minIncrementForUnique(nums []int) int {
	var res int
	length := len(nums)
	record := make(map[int]bool, length)

	for i := 0; i < length; i++ {
		if _, ok := record[nums[i]]; !ok {
			record[nums[i]] = true
		} else {
			num := nums[i]
			for {
				if _, ok := record[num]; !ok {
					record[num] = true
					break
				}
				num++
				res++
			}
		}
	}
	return res
}

// 别人的解法，排序后，维护一个唯一的值value
// 看当前的值到value需要变化多少次
func minIncrementForUniqueII(A []int) int {
	if len(A) == 0 {
		return 0
	}

	sort.Ints(A)
	res, value := 0, A[0]

	for i := 1; i < len(A); i++ {
		if A[i] <= value {
			value++
			res += value - A[i]
		} else {
			value = A[i]
		}
	}
	return res
}

func main() {
	d := []int{3, 2, 1, 2, 1, 7}
	fmt.Println(minIncrementForUniqueII(d))
}
