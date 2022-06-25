package main

import "fmt"

// 一开始想的是双指针做法，但是碰到了反例：aa + ab = aaba
// 双指针一旦移动了，就会产生不可逆的选择，导致条件考虑不全
//func isInterleave(s1 string, s2 string, s3 string) bool {
//	var (
//		index1, index2, index3 int
//		len1, len2, len3       int = len(s1), len(s2), len(s3)
//	)
//	if len1+len2 != len3 {
//		return false
//	}
//	for index3 = 0; index3 < len3; {
//		if index1 < len1 {
//			if s3[index3] != s1[index1] {
//				return false
//			}
//			for index3 < len3 && index1 < len1 && s3[index3] == s1[index1] {
//				index1++
//				index3++
//			}
//		}
//		if index2 < len2 {
//			if s3[index3] != s2[index2] {
//				return false
//			}
//			for index3 < len3 && index2 < len2 && s3[index3] == s2[index2] {
//				index2++
//				index3++
//			}
//		}
//	}
//	return true
//}

func isInterleave(s1 string, s2 string, s3 string) bool {
	n, m, t := len(s1), len(s2), len(s3)
	if (n + m) != t {
		return false
	}
	f := make([][]bool, n+1)
	for i := 0; i <= n; i++ {
		f[i] = make([]bool, m+1)
	}
	f[0][0] = true
	for i := 0; i <= n; i++ {
		for j := 0; j <= m; j++ {
			p := i + j - 1
			if i > 0 {
				f[i][j] = f[i][j] || (f[i-1][j] && s1[i-1] == s3[p])
			}
			if j > 0 {
				f[i][j] = f[i][j] || (f[i][j-1] && s2[j-1] == s3[p])
			}
		}
	}
	return f[n][m]
}

func main() {
	a := "aa"
	b := "ab"
	c := "aaba"
	fmt.Println(isInterleave(a, b, c))
}
