package main

import "fmt"

type Node struct {
	Val  int
	Next *Node
}

// 人形编译器，考虑完所有情况
// 30min
func insert(aNode *Node, x int) *Node {
	tmpNode := &Node{x, nil}
	// 特殊情况，节点数目为0或1
	if aNode == nil {
		tmpNode.Next = tmpNode
		return tmpNode
	} else if aNode.Next == nil {
		aNode.Next = tmpNode
		tmpNode.Next = aNode
		return aNode
	}

	// 常规情况
	prev, next := aNode, aNode.Next
	for {
		// 正常的中间插入
		if x >= prev.Val && x <= next.Val {
			prev.Next = tmpNode
			tmpNode.Next = next
			break
		}
		// 遇到起始点
		if prev.Val > next.Val {
			if x <= prev.Val && x < next.Val || x >= prev.Val {
				prev.Next = tmpNode
				tmpNode.Next = next
				break
			}
		}
		// 最大的末尾插入
		if next.Next == aNode && x >= next.Val {
			next.Next = tmpNode
			tmpNode.Next = aNode
			break
		}
		// 最小的头部插入
		if next == aNode {
			if x < next.Val && x < prev.Val {
				prev.Next = tmpNode
				tmpNode.Next = next
			}
			break
		}
		prev = prev.Next
		next = next.Next
	}
	return aNode
}

// 官方解法
// 官方的解法和我思路一样，但是整体写法简化了很多
func insertII(head *Node, insertVal int) *Node {
	node := &Node{Val: insertVal}
	if head == nil {
		node.Next = node
		return node
	}
	if head.Next == head {
		head.Next = node
		node.Next = head
		return head
	}
	curr, next := head, head.Next
	for next != head {
		if insertVal >= curr.Val && insertVal <= next.Val {
			break
		}
		if curr.Val > next.Val {
			if insertVal > curr.Val || insertVal < next.Val {
				break
			}
		}
		curr = curr.Next
		next = next.Next
	}
	curr.Next = node
	node.Next = next
	return head
}

// 我写的解法——简化版
func insertIII(aNode *Node, x int) *Node {
	tmpNode := &Node{x, nil}
	// 特殊情况，节点数目为0或1
	if aNode == nil {
		tmpNode.Next = tmpNode
		return tmpNode
	} else if aNode.Next == nil {
		aNode.Next = tmpNode
		tmpNode.Next = aNode
		return aNode
	}

	// 常规情况
	prev, next := aNode, aNode.Next
	for {
		// 正常的中间插入
		if x >= prev.Val && x <= next.Val {
			break
		}
		// 遇到起始点
		if prev.Val > next.Val {
			if x <= prev.Val && x < next.Val || x >= prev.Val {
				break
			}
		}
		// 最小的头部插入
		if next == aNode {
			if x < next.Val && x < prev.Val {
			}
			break
		}
		prev = prev.Next
		next = next.Next
	}
	prev.Next = tmpNode
	tmpNode.Next = next
	return aNode
}

func main() {
	a := &Node{3, nil}
	b := &Node{5, nil}
	c := &Node{1, nil}
	a.Next = b
	b.Next = c
	c.Next = a
	res := insert(a, 0)
	fmt.Println(res)
}
