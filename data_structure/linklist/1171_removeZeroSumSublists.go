package main

import "fmt"

//type ListNode struct {
//	Val  int
//	Next *ListNode
//}

// 暴力破解n^2，两个90%

func removeZeroSumSublists(head *ListNode) *ListNode {
	if head == nil {
		return nil
	}
	dummy := &ListNode{-1, head}
	prev := dummy
	low, high := head, head
	var sum int
	for low != nil {
		sum = 0
		for high != nil {
			sum += high.Val
			high = high.Next
			if sum == 0 {
				// 发现和为0，截断中间部分
				prev.Next = high
				low = high
			}
		}
		if low != nil {
			// low往后移动，prev也要跟着一起移动
			low = low.Next
			prev = prev.Next
		}
		high = low
	}
	return dummy.Next
}

func main() {
	d := &ListNode{1, &ListNode{2, &ListNode{3, &ListNode{-3, &ListNode{4, nil}}}}}
	r := removeZeroSumSublists(d)
	fmt.Println(r.Val)
}
