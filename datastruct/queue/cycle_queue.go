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

func (cq *CycleQueue) EnCycleQueue(value int) bool {
	if cq.container == nil || cq.IsFull() {
		return false
	}
	cq.container[cq.rear%len(cq.container)] = value
	cq.rear = cq.rear%len(cq.container) + 1
	cq.size++
	return true
}

func (cq *CycleQueue) DeCycleQueue() (bool, int) {
	if cq.container == nil || cq.IsEmpty() {
		return false, 0
	} else {
		ret := cq.container[cq.front%len(cq.container)]
		cq.front = cq.front%len(cq.container) + 1
		cq.size--
		return true, ret
	}
}

func (cq *CycleQueue) IsEmpty() bool {
	if cq.size == 0 {
		return true
	}
	return false
}

func (cq *CycleQueue) IsFull() bool {
	if cq.size == len(cq.container) {
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
