package main

import "fmt"

func hasCycle(head *ListNode) bool {
	if head == nil {
		return false
	}
	low, high := head, head
	for low != nil && high != nil {
		low = low.Next
		high = high.Next
		if high == nil {
			break
		}
		high = high.Next
		if low == high {
			return true
		}
	}
	return false
}

func main() {
	fmt.Println()
}
