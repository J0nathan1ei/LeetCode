package main

import "fmt"

func getIntersectionNode_(headA, headB *ListNode) *ListNode {
	if headA == nil || headB == nil {
		return nil
	}
	nodeA, nodeB := headA, headB
	var lengthA, lengthB int
	for nodeA != nil {
		lengthA++
		nodeA = nodeA.Next
	}
	for nodeB != nil {
		lengthB++
		nodeB = nodeB.Next
	}
	nodeA, nodeB = headA, headB
	if lengthA > lengthB {
		for i := 0; i < lengthA-lengthB; i++ {
			nodeA = nodeA.Next
		}
	} else {
		for i := 0; i < lengthB-lengthA; i++ {
			nodeB = nodeB.Next
		}
	}

	for nodeA != nil && nodeB != nil {
		if nodeA == nodeB {
			return nodeA
		}
		nodeA = nodeA.Next
		nodeB = nodeB.Next
	}
	return nil
}

func main() {
	fmt.Println()
}
