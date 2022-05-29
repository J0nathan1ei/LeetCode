package main

import "fmt"

func mergeTwoLists(l1 *ListNode, l2 *ListNode) *ListNode {
	if l1 == nil && l2 == nil {
		return nil
	} else if l1 == nil {
		return l2
	} else if l2 == nil {
		return l1
	}

	dummy := &ListNode{0, nil}
	if l1.Val < l2.Val {
		dummy.Next = l1
		l1 = l1.Next
	} else {
		dummy.Next = l2
		l2 = l2.Next
	}
	head := dummy.Next

	for l1 != nil && l2 != nil {
		if l1.Val < l2.Val {
			head.Next = l1
			head = head.Next
			l1 = l1.Next
		} else {
			head.Next = l2
			head = head.Next
			l2 = l2.Next
		}
	}

	if l1 != nil {
		head.Next = l1
	}

	if l2 != nil {
		head.Next = l2
	}

	return dummy.Next
}

func main() {
	fmt.Println()
}
