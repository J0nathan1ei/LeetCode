package main

import "fmt"

//type ListNode struct {
//     Val int
//     Next *ListNode
//}

// 时间复杂度为n^2，空间复杂度为1
//func reorderList(head *ListNode) {
//	if head == nil{
//		return
//	}
//	length := getListLen(head)
//	tmpNode, tmp := head, head
//	left := 0
//	for tmpNode != nil{
//		tmp = tmpNode
//		step := (length - 1 - left) - left
//		if step <= 0{
//			break
//		}
//		for i := 0; i < step;i++{
//			tmp = tmp.Next
//		}
//		tmp.Next = tmpNode.Next
//		tmpNode.Next = tmp
//		tmpNode = tmp.Next
//		left += 1
//	}
//	tmp.Next = nil
//}
//
//func getListLen(head *ListNode)int{
//	if head == nil{
//		return 0
//	}
//	var res int = 0
//	for tmp := head; tmp!= nil;tmp = tmp.Next{
//		res++
//	}
//	return res
//}

func reorderList(head *ListNode) {
	if head == nil {
		return
	}

	var nodes []*ListNode
	tmpHead := head
	for ; tmpHead != nil; tmpHead = tmpHead.Next {
		nodes = append(nodes, tmpHead)
	}
	length := len(nodes)
	mid := (length - 1) / 2
	for i := 0; i <= mid; i++ {
		j := length - 1 - i
		if j-i <= 1 {
			nodes[j].Next = nil
			break
		}
		nodes[j].Next = nodes[i].Next
		nodes[i].Next = nodes[j]
	}
}

func main() {
	d := &ListNode{1, &ListNode{2, &ListNode{3, &ListNode{4, nil}}}}
	reorderList(d)
	fmt.Println(d.Val)
}
