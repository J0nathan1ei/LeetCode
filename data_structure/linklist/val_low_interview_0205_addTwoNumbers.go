package main

import "fmt"

type ListNode struct {
	Val  int
	Next *ListNode
}

func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	head1 := l1
	var num, cnt int
	for l1 != nil && l2 != nil {
		num = l1.Val + l2.Val
		if cnt > 0 {
			num += cnt
			cnt = 0
		}
		if num >= 10 {
			cnt = 1
			num -= 10
		}
		l1.Val = num

		l1 = l1.Next
		l2 = l2.Next
	}

	if l1 == nil && l2 != nil {
		for l2 != nil {
			num = l2.Val
			if cnt > 0 {
				num = l2.Val + cnt
			}
			if num >= 10 {
				cnt = 1
				num -= 10
			}
			l2.Val = num
			l2 = l2.Next
		}
	} else if l2 == nil && l1 != nil {
		for l1 != nil {
			num = l1.Val
			if cnt > 0 {
				num = l1.Val + cnt
			}
			if num >= 10 {
				cnt = 1
				num -= 10
			}
			l1.Val = num
			l1 = l1.Next
		}
	}
	if cnt > 0 {
		tmp := head1
		for tmp.Next != nil {
			tmp = tmp.Next
		}
		tmp.Next = &ListNode{1, nil}
	}
	return head1
}

func main() {
	l1 := &ListNode{5, nil}
	l2 := &ListNode{5, nil}
	d := addTwoNumbers(l1, l2)
	fmt.Println(d.Val)
}
