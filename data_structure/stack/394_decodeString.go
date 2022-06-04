package main

import (
	"fmt"
	"strings"
)

// 尝试1.尝试使用双栈，将数字和字母区分开，
// 出现的问题是：数字和字母无法一一对应，而且增大了耦合的复杂度
func decodeStringdoubleStack(s string) string {
	var res string
	var stackNum []int
	var stackChars []string
	length := len(s)
	for i := 0; i < length; i++ {
		if isNum(s[i]) {
			tmpNum := ""
			for isNum(s[i]) {
				tmpNum += string(s[i])
				i++
			}
			stackNum = append(stackNum, stringToNum(tmpNum))
			i--
		} else if s[i] == '[' {
			continue
		} else if isChar(s[i]) {
			tmpChars := ""
			for isChar(s[i]) {
				tmpChars += string(s[i])
				i++
			}
			stackChars = append(stackChars, tmpChars)
			i--
		} else if s[i] == ']' {
			subString := ""
			for j := 0; j < stackNum[len(stackNum)-1]; j++ {
				subString += stackChars[len(stackChars)-1]
			}
			// 出栈
			stackChars = stackChars[:len(stackChars)-1]
			stackNum = stackNum[:len(stackNum)-1]
			if len(stackChars) == 0 {
				res = subString
				break
			} else {
				stackChars[len(stackChars)-1] += subString
			}
		}
	}
	return res
}

// 尝试2. 尝试使用strings.Split将原字符串按“[”来分隔，
// 出现的问题是：有些出现一次的字符串，前面没有“[”，会导致混淆错误

// 只能使用单栈，因为可以维护数字和字母的先后顺序
func decodeString(s string) string {
	var stack []string
	var res string
	length := len(s)
	for i := 0; i < length; i++ {
		if isNum(s[i]) {
			tmpNum := ""
			for i < length && isNum(s[i]) {
				tmpNum += string(s[i])
				i++
			}
			i--
			stack = append(stack, tmpNum)
		} else if s[i] == '[' {
			stack = append(stack, "[")
			continue
		} else if isChar(s[i]) {
			tmpStr := ""
			for i < length && isChar(s[i]) {
				tmpStr += string(s[i])
				i++
			}
			i--
			stack = append(stack, tmpStr)
		} else {
			var repStr string
			for stack[len(stack)-1] != "[" {
				repStr = stack[len(stack)-1] + repStr
				stack = stack[:len(stack)-1]
			}
			stack = stack[:len(stack)-1]
			repTimes := stringToNum(stack[len(stack)-1])
			stack[len(stack)-1] = ""
			for j := 0; j < repTimes; j++ {
				stack[len(stack)-1] += repStr
			}
		}
	}
	for i := 0; i < len(stack); i++ {
		res += stack[i]
	}

	return res
}

func isNum(s byte) bool {
	if s >= '0' && s <= '9' {
		return true
	}
	return false
}

func stringToNum(s string) int {
	var res int
	for i := 0; i < len(s); i++ {
		res = res*10 + int(s[i]-'0')
	}
	return res
}

func isChar(s byte) bool {
	if s >= 'a' && s <= 'z' || s >= 'A' && s <= 'Z' {
		return true
	}
	return false
}

func main() {
	d := strings.Split("3[a2[c3[d]]]", "[")
	fmt.Println(d)
	//fmt.Println(decodeString("abc3[cd]xyz"))
	//fmt.Println(decodeString("3[a]2[bc]"))
	fmt.Println(decodeString("3[z]2[2[y]pq4[2[jk]e1[f]]]ef"))
}
