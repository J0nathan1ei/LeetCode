package main

import "fmt"

/*
基本的链表反转法：
维护三个节点 prev、cur和next
*/
func reverseList(head *ListNode) *ListNode {
	if head == nil {
		return head
	}
	var prev, cur, next *ListNode
	cur = head
	for cur != nil {
		next = cur.Next
		cur.Next = prev
		prev = cur
		cur = next
	}
	return prev
}

func main() {
	fmt.Println()
}
