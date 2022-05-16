package main

import (
	"container/list"
)
/*
基本思想是使用一个队列，每次更新的值放到队首，淘汰时从队尾淘汰
*/

// TODO
type LRUCache struct {
	queue list.List
	m map[int]int
}


func Constructor(capacity int) LRUCache {
	res := &LRUCache{}
	res.m = make(map[int]int, capacity)
}


func (this *LRUCache) Get(key int) int {

}


func (this *LRUCache) Put(key int, value int)  {

}

func main(){

}
/**
 * Your LRUCache object will be instantiated and called as such:
 * obj := Constructor(capacity);
 * param_1 := obj.Get(key);
 * obj.Put(key,value);
 */