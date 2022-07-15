package main

import (
	"fmt"
	"sort"
)

//初步思路是排一次序，然后按顺序往出取，取一个就删除一个元素
// 这样时间复杂度是NlogN，空间复杂度1
// 数组处理的时候，有很多边界条件细节
// time: 30min+
func isNStraightHand(hand []int, groupSize int) bool {
	// 不能整除则return false
	if len(hand)%groupSize != 0 || len(hand) == 0 {
		return false
	}

	// 排序，遍历查找
	sort.Ints(hand)
	preNum := hand[0]
	hand = hand[1:]
	cnt := 1
	for i := 0; i < len(hand); i++ {
		// 找到符合条件的元素，则删除此元素，寻找下一个
		if hand[i] == preNum+1 {
			hand = append(hand[:i], hand[i+1:]...)
			preNum++
			cnt++
			i--
		}
		// 已经找到groupSize大小的组合，继续从头开始，找下一组
		if cnt == groupSize && len(hand) > 0 {
			i = -1
			preNum = hand[0]
			hand = hand[1:]
			cnt = 1
		} else if i == len(hand)-1 && cnt < groupSize {
			// 到达末尾也没有找到groupSize大小的数组
			return false
		}
	}
	return true
}
func main() {
	d := []int{1, 2, 3, 4, 5, 4}
	size := 3
	fmt.Println(isNStraightHand(d, size))
}
