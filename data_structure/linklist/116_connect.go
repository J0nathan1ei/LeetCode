package main

type Node struct {
	Val   int
	Left  *Node
	Right *Node
	Next  *Node
}

func connect(root *Node) *Node {
	if root == nil {
		return nil
	}
	var tmpList []*Node
	tmpRoot := root
	tmpList = append(tmpList, tmpRoot)

	for len(tmpList) > 0 {
		length := len(tmpList)
		for i := 0; i < length; i++ {
			if i == length-1 {
				tmpList[i].Next = nil
			} else {
				tmpList[i].Next = tmpList[i+1]
			}

			if tmpList[i].Left != nil {
				tmpList = append(tmpList, tmpList[i].Left)
			}
			if tmpList[i].Right != nil {
				tmpList = append(tmpList, tmpList[i].Right)
			}
		}
		tmpList = tmpList[length:]
	}
	return root
}
