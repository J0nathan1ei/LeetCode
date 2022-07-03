package main

import "fmt"

// easy
// 双指针，3min解决
func middleNode(head *ListNode) *ListNode {
	if head == nil {
		return nil
	}
	dummy := &ListNode{-1, head}
	low, high := dummy, dummy
	for {
		if high == nil {
			return low
		} else if high.Next == nil {
			return low.Next
		}

		high = high.Next.Next
		low = low.Next
	}
}
func main() {
	fmt.Println()
}
