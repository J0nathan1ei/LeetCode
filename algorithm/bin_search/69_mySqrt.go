package main

import "fmt"

// 二分查找法实现的mySqrt
// 注意循环退出的条件，在平方刚好大于x时退出，就需要mid--
func mySqrt(x int) int {
	if x == 0 || x == 1 {
		return x
	}
	low, high := 1, x
	mid, tmpRes := 0, 0
	for low <= high {
		mid = low + (high-low)/2
		tmpRes = mid * mid
		if tmpRes == x {
			return mid
		} else if tmpRes > x {
			high = mid - 1
			// 这里为什么要mid--，因为循环退出的条件也可能是平方刚好大于x
			// 所以需要减一
			mid--
		} else {
			low = mid + 1
		}
	}
	return mid
}

func main() {
	fmt.Println(mySqrt(999))
}
