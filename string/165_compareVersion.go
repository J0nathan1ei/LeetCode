package main

import (
	"fmt"
	"strconv"
	"strings"
)

// 字符串比较，split+Atoi即可
// 11min 20s完成
func compareVersion(version1 string, version2 string) int {
	var (
		strs1      = strings.Split(version1, ".")
		strs2      = strings.Split(version2, ".")
		len1, len2 = len(strs1), len(strs2)
		i          int
	)

	for i = 0; ; i++ {
		if i == len1 || i == len2 {
			break
		}
		num1, _ := strconv.Atoi(strs1[i])
		num2, _ := strconv.Atoi(strs2[i])
		if num1 > num2 {
			return 1
		} else if num1 < num2 {
			return -1
		}
	}

	if len1 > len2 {
		for ; i < len1; i++ {
			num1, _ := strconv.Atoi(strs1[i])
			if num1 > 0 {
				return 1
			}
		}
	} else if len1 < len2 {
		for ; i < len2; i++ {
			num2, _ := strconv.Atoi(strs2[i])
			if num2 > 0 {
				return -1
			}
		}
	}
	return 0
}
func main() {
	fmt.Println(compareVersion("0.1", "1.0.0"))
}
