package main

//type ListNode struct {
//	Val  int
//	Next *ListNode
//}

func getIntersectionNode(headA, headB *ListNode) *ListNode {
	if headA == nil || headB == nil {
		return nil
	}
	tmpHeadA, tmpHeadB := headA, headB
	var lenA, lenB int = 1, 1
	for ; tmpHeadA != nil; tmpHeadA = tmpHeadA.Next {
		lenA++
	}
	for ; tmpHeadB != nil; tmpHeadB = tmpHeadB.Next {
		lenB++
	}
	tmpHeadA = headA
	tmpHeadB = headB
	if lenA > lenB {
		for i := 0; i < lenA-lenB; i++ {
			tmpHeadA = tmpHeadA.Next
		}
	} else if lenB > lenA {
		for i := 0; i < lenB-lenA; i++ {
			tmpHeadB = tmpHeadB.Next
		}
	}

	for tmpHeadA != nil && tmpHeadB != nil {
		if tmpHeadA == tmpHeadB {
			return tmpHeadA
		}
		tmpHeadA = tmpHeadA.Next
		tmpHeadB = tmpHeadB.Next
	}
	return nil
}

func main() {

}
