package 堆排序

import (
	"fmt"
)

/*
	parent=(i-1)/2
	child1=2i+1
	child2=2i+2
*/


type Heap struct {
	data []int
	len  int
	cap  int
}

// 构造堆过程
func NewHeap(length int) *Heap {
	return &Heap{
		data: make([]int, length),
		cap:  length,
		len:  0,
	}
}
func (h *Heap) Insert(element int) bool {
	if h.len+1 > h.cap {
		return false
	}
	h.shiftUp(element)
	h.len = h.len + 1
	return true
}
func (h *Heap) Remove() (int, bool) {
	if h.len <= 0 {
		return 0, false
	}
	var ret = h.shiftDown()
	h.len--
	return ret, true
}

// 注意: 元素数量个下标是差1 的关系
func (h *Heap) shiftUp(element int) {
	// 添加到最后
	h.data[h.len] = element
	parent := h.len
	// 然后循环比较
	for ; h.len > 1; {

		if h.data[parent] > h.data[(parent-1)/2] {
			h.data[parent], h.data[(parent-1)/2] = h.data[(parent-1)/2], h.data[parent]
			parent = (parent - 1) / 2
			if parent == 0 {
				break
			}
			continue
		}
		break
	}
}

func (h *Heap) shiftDown() int {
	ret := h.data[0]
	h.data[0] = h.data[h.len-1]
	// 数组里面还有len-1个元素 最大的下标值是 h.len-2
	h.data[h.len-1] = 0
	for i := 0; i <= h.len-2; {
		// 先寻找左右两个子节点最大值
		var max int
		var maxPosition int
		//  此时应百世
		if 2*i+1 > h.len-2 {

		}
		if h.data[i*2+1] > h.data[i*2+2] {
			max = h.data[i*2+1]
			maxPosition = i*2 + 1
		} else {
			max = h.data[i*2+2]
			maxPosition = i*2 + 2
		}
		if h.data[i] >= max {
			break
		}
		h.data[i], h.data[maxPosition] = h.data[maxPosition], h.data[i]
		i = maxPosition
	}
	fmt.Println(h.data)
	return ret
}
