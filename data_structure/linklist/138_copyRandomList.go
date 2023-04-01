package main

type Node struct {
	Val    int
	Next   *Node
	Random *Node
}

// 做法比较中规中矩， 缺点是要申请数组来维护random的关系
func copyRandomList(head *Node) *Node {
	if head == nil {
		return nil
	}
	// 第一遍遍历，获取链表的 节点地址-序号 映射
	dictNodeToIndex := make(map[*Node]int)
	headIndex := head
	var cnt int
	for ; headIndex != nil; headIndex = headIndex.Next {
		dictNodeToIndex[headIndex] = cnt
		cnt++
	}
	dictNodeToIndex[nil] = cnt

	// 第二遍遍历，获取链表的 random 序号-序号 映射
	dictIndexToRandom := make(map[int]int)
	for headIndex = head; headIndex != nil; headIndex = headIndex.Next {
		dictIndexToRandom[dictNodeToIndex[headIndex]] = dictNodeToIndex[headIndex.Random]
	}

	newHead := &Node{
		Val:    head.Val,
		Next:   nil,
		Random: nil,
	}
	newHeadI := newHead
	dictNewIndexToNode := make(map[int]*Node)
	cnt = 0
	// 新链表的第一次遍历，建立链表，同时获取新链表的 序号-节点地址 映射
	for headIndex = head.Next; headIndex != nil; headIndex = headIndex.Next {
		newHeadI.Next = &Node{
			Val:    headIndex.Val,
			Next:   nil,
			Random: nil,
		}
		dictNewIndexToNode[cnt] = newHeadI
		cnt++
		newHeadI = newHeadI.Next
	}
	dictNewIndexToNode[cnt] = newHeadI
	cnt++
	dictNewIndexToNode[cnt] = nil

	// 新链表的第二次遍历，利用之前的 序号-序号 映射 和 序号-节点地址 映射 建立random的关系
	cnt = 0
	for newHeadI = newHead; newHeadI != nil; newHeadI = newHeadI.Next {
		newHeadI.Random = dictNewIndexToNode[dictIndexToRandom[cnt]]
		cnt++
	}

	return newHead
}
