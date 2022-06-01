package main

import "fmt"

//type ListNode struct {
//	Val  int
//	Next *ListNode
//}

func reverseBetween(head *ListNode, left int, right int) *ListNode {
	if left == right {
		return head
	}
	tmpHead := &ListNode{-1, head}
	before := tmpHead
	cur, prev, next := head, head, head
	for i := 1; i < left; i++ {
		before = before.Next
	}
	for i := 1; i < left; i++ {
		cur = cur.Next
	}
	for i := 1; i <= right; i++ {
		prev = prev.Next
	}
	for i := 0; i <= right-left; i++ {
		next = cur.Next
		cur.Next = prev
		prev = cur
		before.Next = cur // 放在这里，不是很美观
		cur = next
	}
	//before.Next = cur 一开始把before的赋值写在这里，不对，cur变成了应该指向的下一个节点

	return tmpHead.Next
}

func main() {
	d := &ListNode{1, &ListNode{2, &ListNode{3, &ListNode{4, &ListNode{5, nil}}}}}
	res := reverseBetween(d, 2, 4)
	fmt.Println(res)
}
