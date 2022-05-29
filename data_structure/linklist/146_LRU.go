package main

import "fmt"

// 经典高频算法实现

type LRUCache struct {
	m          map[int]*LruNode
	head, tail *LruNode
	size, cap  int
}

type LruNode struct {
	key, val   int
	prev, next *LruNode
}

func initLruNode(key, val int) *LruNode {
	return &LruNode{
		key, val,
		nil, nil,
	}
}

// 删除单个节点
func (this *LRUCache) removeNode(node *LruNode) {
	node.next.prev = node.prev
	node.prev.next = node.next
}

func (this *LRUCache) addToHead(node *LruNode) {
	// 处理插入后的当前节点
	node.next = this.head.next
	node.prev = this.head
	// 处理插入后的节点后面的那个节点
	node.next.prev = node
	// 处理头结点
	this.head.next = node
}

func (this *LRUCache) moveToHead(node *LruNode) {
	this.removeNode(node)
	this.addToHead(node)
}

func (this *LRUCache) deleteTail() {
	node := this.tail.prev
	if node != this.head {
		node.prev.next = node.next
		node.next.prev = node.prev
		delete(this.m, node.key)
	}
	if this.size > 1 {
		this.size--
	}
}

func Constructor(capacity int) LRUCache {
	l := LRUCache{
		m:    map[int]*LruNode{},
		head: initLruNode(0, 0),
		tail: initLruNode(0, 0),
		size: 0,
		cap:  capacity,
	}
	l.head.next = l.tail
	l.tail.prev = l.head
	return l
}

func (this *LRUCache) Get(key int) int {
	if _, ok := this.m[key]; !ok {
		return -1
	} else {
		node := this.m[key]
		this.moveToHead(node)
		return node.val
	}
}

func (this *LRUCache) Put(key int, value int) {
	if _, ok := this.m[key]; !ok {
		node := initLruNode(key, value)
		this.m[key] = node
		this.addToHead(node)
		this.size++
		if this.size > this.cap {
			this.deleteTail()
		}
	} else {
		node := this.m[key]
		node.val = value
		this.moveToHead(node)
	}
}

func main() {
	d := Constructor(2)
	d.Put(2, 1)
	d.Put(2, 2)
	fmt.Println(d.Get(2))
	d.Put(1, 1)
	d.Put(4, 1)
	fmt.Println(d.Get(2))
	fmt.Println()
}
