package main

import (
	"fmt"
)

// 利用go的map的无序性
// 时间：30min
type PhoneDirectory struct {
	m map[int]bool
}

func Constructor(maxNumbers int) PhoneDirectory {
	res := PhoneDirectory{
		m: make(map[int]bool, maxNumbers),
	}
	for i := 0; i < maxNumbers; i++ {
		res.m[i] = true
	}
	return res
}

func (this *PhoneDirectory) Get() int {
	// map中不存在元素则返回-1
	if len(this.m) == 0 {
		return -1
	}
	// 标准情况，检查已分配的号码
	var res int
	for key := range this.m {
		res = key
		delete(this.m, key)
		break
	}
	return res
}

func (this *PhoneDirectory) Check(number int) bool {
	if _, ok := this.m[number]; ok {
		return true
	}
	return false
}

func (this *PhoneDirectory) Release(number int) {
	if _, ok := this.m[number]; !ok {
		this.m[number] = true
	}
}

/**
 * Your PhoneDirectory object will be instantiated and called as such:
 * obj := Constructor(maxNumbers);
 * param_1 := obj.Get();
 * param_2 := obj.Check(number);
 * obj.Release(number);
 */
func main() {
	d := Constructor(1)
	a := d.Get()
	b := d.Get()
	fmt.Println(a, b)
	fmt.Println(d.Check(1))
	d.Release(1)
	fmt.Println(d.Check(1))

}
