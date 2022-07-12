package main

import (
	"fmt"
	"strconv"
	"strings"
)

// 基本的IP地址判断模拟
// 细节比较多，尤其是IPv4的判断
// 时间：30min+
func validIPAddress(queryIP string) string {
	pointCnt := strings.Count(queryIP, ".")
	colonCnt := strings.Count(queryIP, ":")
	if pointCnt > 0 && colonCnt == 0 && judgeIPV4(queryIP) {
		return "IPv4"
	}
	if colonCnt > 0 && pointCnt == 0 && judgeIPV6(queryIP) {
		return "IPv6"
	}
	return "Neither"
}

func judgeIPV4(s string) bool {
	numSplit := strings.Split(s, ".")
	if len(numSplit) != 4 {
		return false
	}

	for i := 0; i < 4; i++ {
		// 可以一个元素是0，但是不能是两个以上元素以0开头
		if len(numSplit[i]) > 1 && numSplit[i][0] == '0' {
			return false
		}
		// 转换错误，转换后的数字超出范围，是违规情况
		num, err := strconv.Atoi(numSplit[i])
		if err != nil || num < 0 || num > 255 {
			return false
		}
	}
	return true
}

func judgeIPV6(s string) bool {
	addressSplit := strings.Split(s, ":")
	if len(addressSplit) != 8 {
		return false
	}
	for i := 0; i < 8; i++ {
		if !judgeIPV6Unit(addressSplit[i]) {
			return false
		}
	}
	return true
}

func judgeIPV6Unit(s string) bool {
	if len(s) == 0 || len(s) > 4 {
		return false
	}
	s = strings.ToLower(s)
	for _, val := range s {
		if (val < '0' || val > '9') && (val < 'a' || val > 'f') {
			return false
		}
	}
	return true
}

func main() {
	d := "192.0.0.1"
	fmt.Println(validIPAddress(d))
}
