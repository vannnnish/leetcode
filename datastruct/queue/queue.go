package queue

type Queue struct {
	container []int
	// 元素位置标记
	index int
	size  int
}

func NewQueue(size int) *Queue {
	return &Queue{
		container: make([]int, size),
		index:     0,
		size:      0,
	}
}

func (q *Queue) IsEmpty() bool {
	return q.size == 0
}

func (q *Queue) IsFull() bool {
	return q.size == len(q.container)
}

func (q *Queue) EnQueue(a int) bool {
	if q.IsFull() {
		return false
	}
	q.container[q.index] = a
	q.index++
	q.size++
	return true
}

func (q *Queue) DeQueue() (bool, int) {
	if q.IsEmpty() {
		return false, 0
	}
	ret := q.container[0]
	// 移动元素位置
	for i := 1; i < q.size; i++ {
		q.container[i-1] = q.container[i]
	}
	q.index--
	q.size--
	return true, ret
}
