package main

import "fmt"

//type ListNode struct {
//	Val  int
//	Next *ListNode
//}

func rotateRight(head *ListNode, k int) *ListNode {
	if head == nil || head.Next == nil || k == 0 {
		return head
	}
	dummy := &ListNode{-1, head}
	// 获取链表长度
	var length int = 1
	tail := head
	for ; tail.Next != nil; tail = tail.Next {
		length++
	}

	// 取余，删除多余的移动次数
	k = k % length
	// 仅当k != 0时，才移动节点
	if k != 0 {
		mov := head
		for i := 0; i < length-k-1; i++ {
			mov = mov.Next
		}

		// 移动链表结构
		dummy.Next = mov.Next
		mov.Next = nil
		tail.Next = head
	}

	return dummy.Next
}
func main() {
	d := &ListNode{1, &ListNode{2, nil}}
	r := rotateRight(d, 2)
	fmt.Println(r.Val)
}
