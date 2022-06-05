package main

import "fmt"

// 写法很长
//func calculate(s string) int {
//	var stackNum []int
//	var stackCalcu []byte
//	for i := 0; i < len(s); i++ {
//		if s[i] == ' ' {
//			continue
//		} else if isNum(s[i]) {
//			num := getNum(s, &i)
//			stackNum = append(stackNum, num)
//		} else if s[i] == '+' || s[i] == '-' {
//			stackCalcu = append(stackCalcu, s[i])
//		} else if s[i] == '*' {
//			i++
//			num := getNum(s, &i)
//			stackNum[len(stackNum)-1] = stackNum[len(stackNum)-1] * num
//		} else if s[i] == '/' {
//			i++
//			num := getNum(s, &i)
//			stackNum[len(stackNum)-1] = stackNum[len(stackNum)-1] / num
//		}
//	}
//	for i := 0; i < len(stackCalcu); i++ {
//		num1, num2 := stackNum[0], stackNum[1]
//		stackNum = stackNum[1:]
//		if stackCalcu[i] == '+' {
//			stackNum[0] = num1 + num2
//		} else {
//			stackNum[0] = num1 - num2
//		}
//	}
//	return stackNum[0]
//}
//
//func getNum(s string, index *int) int {
//	tmpNum := ""
//	for ; *index < len(s) && (isNum(s[*index]) || s[*index] == ' '); *index++ {
//		if isNum(s[*index]) {
//			tmpNum += string(s[*index])
//		}
//	}
//	*index--
//	return stringToNum(tmpNum)
//}
//
//func stringToNum(s string) int {
//	var res int
//	for i := 0; i < len(s); i++ {
//		res = res*10 + int(s[i]-'0')
//	}
//	return res
//}
//
//func isNum(s byte) bool {
//	if s >= '0' && s <= '9' {
//		return true
//	}
//	return false
//}

// 值得学习的写法，长度是我的一半
func calculate(s string) (ans int) {
	stack := []int{}
	preSign := '+'
	num := 0
	for i, ch := range s {
		isDigit := '0' <= ch && ch <= '9'
		if isDigit {
			num = num*10 + int(ch-'0')
		}
		if !isDigit && ch != ' ' || i == len(s)-1 {
			switch preSign {
			case '+':
				stack = append(stack, num)
			case '-':
				stack = append(stack, -num)
			case '*':
				stack[len(stack)-1] *= num
			default:
				stack[len(stack)-1] /= num
			}
			preSign = ch
			num = 0
		}
	}
	for _, v := range stack {
		ans += v
	}
	return
}

func main() {
	fmt.Println(calculate("0-2147483647"))
	fmt.Println(calculate("1+1-1"))
}
