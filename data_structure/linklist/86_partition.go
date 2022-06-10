package main

import "fmt"

//type ListNode struct {
//	Val  int
//	Next *ListNode
//}

func partition(head *ListNode, x int) *ListNode {
	if head == nil {
		return nil
	}
	var lowQueue []*ListNode
	var highQueue []*ListNode

	dummy := &ListNode{-1, head}
	cur := head
	for ; cur != nil; cur = cur.Next {
		if cur.Val < x {
			lowQueue = append(lowQueue, cur)
		} else {
			highQueue = append(highQueue, cur)
		}
	}
	cur = dummy
	for i := 0; i < len(lowQueue); i++ {
		cur.Next = lowQueue[i]
		cur = cur.Next
	}
	for i := 0; i < len(highQueue); i++ {
		cur.Next = highQueue[i]
		cur = cur.Next
	}
	// 这里不置nil的话，会有环出现，导致无限分配内存
	cur.Next = nil

	return dummy.Next
}

func main() {
	d := &ListNode{1, &ListNode{4, &ListNode{3, &ListNode{
		2, &ListNode{5, &ListNode{2, nil}}}}}}
	r := partition(d, 3)
	fmt.Println(r.Val)
}
