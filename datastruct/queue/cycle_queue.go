package queue

import (
	"fmt"
)

/*
	不用移动元素的办法
*/
type CycleQueue struct {
	container []int
	front     int
	rear      int
	size      int
}

func NewCycleQueue(k int) *CycleQueue {
	return &CycleQueue{
		make([]int, k),
		0,
		0,
		0,
	}
}

func (this *CycleQueue) EnCycleQueue(value int) bool {
	if this.container == nil || this.IsFull() {
		return false
	}
	this.container[this.rear%len(this.container)] = value
	this.rear = this.rear%len(this.container) + 1
	this.size++
	return true
}

func (this *CycleQueue) DeCycleQueue() (bool, int) {
	if this.container == nil || this.IsEmpty() {
		return false, 0
	} else {
		ret := this.container[this.front%len(this.container)]
		this.front = this.front%len(this.container) + 1
		this.size--
		return true, ret
	}
}

func (this *CycleQueue) IsEmpty() bool {
	if this.size == 0 {
		return true
	}
	return false
}

func (this *CycleQueue) IsFull() bool {
	if this.size == len(this.container) {
		return true
	}
	return false
}

func main() {
	CycleQueue := NewCycleQueue(5)
	// 循环3次，每次添加5个元素，再出队三个元素
	for i := 0; i < 3; i++ {
		for j := 0; j < 5; j++ {
			CycleQueue.EnCycleQueue(j)
		}
		for k := 0; k < 3; k++ {
			fmt.Println(CycleQueue.DeCycleQueue())
		}
	}
}
