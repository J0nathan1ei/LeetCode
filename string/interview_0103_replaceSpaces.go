package main

import "fmt"

// "假定该字符串尾部有足够的空间存放新增字符"，实际上字符串后面的空格还有可能给多

func replaceSpaces(S string, length int) string {
	var (
		low  = length - 1
		high = len(S) - 1
		res  = make([]byte, len(S))
	)
	for low >= 0 && high >= 0 {
		if S[low] != ' ' {
			res[high] = S[low]
			high--
		} else {
			res[high] = '0'
			res[high-1] = '2'
			res[high-2] = '%'
			high -= 3
		}
		low--
	}
	return string(res[high+1:])
}

func main() {
	d := "ds sdfs afs sdfa dfssf asdf             "
	fmt.Println(replaceSpaces(d, 27))
}
