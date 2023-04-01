package main

import "fmt"

//type ListNode struct {
//	Val  int
//	Next *ListNode
//}

func oddEvenList(head *ListNode) *ListNode {
	if head == nil {
		return nil
	}
	oddHead := head
	evenHead := head.Next
	even := evenHead
	for even != nil && even.Next != nil {
		oddHead.Next = even.Next
		oddHead = even.Next
		even.Next = oddHead.Next
		even = oddHead.Next
	}
	oddHead.Next = evenHead
	return head
}

func main() {
	d := &ListNode{1,
		&ListNode{2,
			&ListNode{3,
				&ListNode{4,
					&ListNode{5, nil}},
			}}}
	d = oddEvenList(d)
	for i := d; i != nil; i = i.Next {
		fmt.Println(i.Val)
	}
}
