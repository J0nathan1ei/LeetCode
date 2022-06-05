package main

import "fmt"

//type ListNode struct {
//	Val  int
//	Next *ListNode
//}

func deleteDuplicatesII(head *ListNode) *ListNode {
	if head == nil {
		return nil
	}
	dummy := &ListNode{-1, head}
	prev, low, high := dummy, head, head.Next
	for low != nil && high != nil {
		if high.Val == low.Val {
			for high != nil && high.Val == low.Val {
				high = high.Next
			}
			// 跳过中间重复节点
			prev.Next = high
			low = high
			if high != nil {
				high = high.Next
			}
		} else {
			low = low.Next
			high = high.Next
			prev = prev.Next
		}
	}
	return dummy.Next
}

func main() {
	d := &ListNode{1, &ListNode{1, nil}}
	r := deleteDuplicatesII(d)
	fmt.Println(r.Val)
}
