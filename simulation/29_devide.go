package main

import (
	"fmt"
	"math"
)

/*基本想法是用移位运算，加减法来逐步逼近结果
在Int32的边界范围卡了几次
*/

func divide(dividend int, divisor int) int {
	var diff bool
	if divisor*dividend > 0 {
		diff = true
	}
	dividend, divisor = abs(dividend), abs(divisor)

	var res int = 0
	tmpDividend, tmpDivisor := dividend, divisor
	for tmpDividend-tmpDivisor > 0 {
		i := 1
		for tmpDividend > tmpDivisor {
			tmpDivisor <<= 1
			i <<= 1
		}
		tmpDivisor >>= 1
		i >>= 1
		res += i
		tmpDividend = tmpDividend - tmpDivisor
		tmpDivisor = divisor
	}
	if tmpDividend == tmpDivisor {
		res += 1
	}

	if !diff {
		res *= -1
	}

	if res < math.MinInt32 {
		return math.MinInt32
	} else if res > math.MaxInt32 {
		return math.MaxInt32
	}
	return res
}

func abs(x int) int {
	if x >= 0 {
		return x
	}
	return -x
}

func main() {
	fmt.Println(divide(-2147483648, -1))
}
