package main

import "fmt"

type CQueue struct {
	stack1, stack2 *[]int
}

func Constructor() CQueue {
	return CQueue{
		stack1: new([]int),
		stack2: new([]int),
	}
}

func (this *CQueue) AppendTail(value int) {
	*this.stack2 = append(*this.stack2, value)
	return
}

func (this *CQueue) DeleteHead() int {
	var ret int
	len1 := len(*this.stack1)
	len2 := len(*this.stack2)
	if len1 > 0 {
		ret := (*this.stack1)[len1-1]
		*this.stack1 = (*this.stack1)[:len1-1]
		return ret
	}

	// 若无元素则返回-1
	if len(*this.stack2) == 0 {
		return -1
	}
	// 从stack2向stack1入栈，入栈后清掉stack2内容
	ret = (*this.stack2)[0]
	for i := len2 - 1; i >= 1; i-- {
		*this.stack1 = append(*this.stack1, (*this.stack2)[i])
	}
	*this.stack2 = []int{}

	return ret
}

func main() {
	demo := Constructor()
	demo.AppendTail(1)
	demo.AppendTail(2)
	demo.AppendTail(3)
	fmt.Println(demo.DeleteHead())
	fmt.Println(demo.DeleteHead())
	demo.AppendTail(4)
	fmt.Println(demo.DeleteHead())
	demo.AppendTail(5)
	fmt.Println(demo.DeleteHead())
	fmt.Println(demo.DeleteHead())

}
