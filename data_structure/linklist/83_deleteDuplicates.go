package main

import "fmt"

func deleteDuplicates(head *ListNode) *ListNode {
	if head == nil {
		return nil
	}
	prev, cur := head, head.Next
	for cur != nil {
		if cur.Val == prev.Val {
			prev.Next = cur.Next
			cur = cur.Next
		}
		if prev.Next != nil && prev.Next.Val != prev.Val {
			prev = prev.Next
		}
	}
	return head
}

func main() {
	fmt.Println()
}
