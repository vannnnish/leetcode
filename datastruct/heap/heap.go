package heap

type MaxHeap struct {
	arr []int
}

func (heap *MaxHeap) Size() int {
	return len(heap.arr)
}

func (heap *MaxHeap) IsEmpty() bool {
	return len(heap.arr) == 0
}

func (heap *MaxHeap) parent(index int) int {
	if index <= 0 {
		panic("0节点没有父节点")
	}
	return (index - 1) / 2
}

func (heap *MaxHeap) leftChild(index int) int {
	return index*2 + 1
}

func (heap *MaxHeap) rightChild(index int) int {
	return index*2 + 2
}

func (heap *MaxHeap) FindMax() int {
	if heap.Size() == 0 {
		panic("堆是空的，不能查询了")
	}
	return heap.arr[0]
}

func (heap *MaxHeap) Add(e int) {
	heap.arr = append(heap.arr, e)
	heap.shiftUp(heap.Size() - 1)
}

func (heap *MaxHeap) shiftUp(size int) {
	// 插入的元素和父节点元素相比较，如果比他大，那么会交换两点的元素，并继续比较
	for size > 0 && heap.arr[heap.parent(size)] < heap.arr[size] {
		heap.arr[heap.parent(size)], heap.arr[size] = heap.arr[size], heap.arr[heap.parent(size)]
		size = heap.parent(size)
	}
}

func (heap *MaxHeap) RemoveMax() int {
	var ret = heap.FindMax()
	heap.arr[0], heap.arr[heap.Size()-1] = heap.arr[heap.Size()-1], heap.arr[0]
	heap.arr = heap.arr[0 : heap.Size()-1]
	heap.shiftDown()
	return ret
}

func (heap *MaxHeap) shiftDown() {
	var (
		j int
		// 最大堆删除的元素只能是第0个
		index = 0
	)
	// 此操作是为了避免一些极限情况的出现
	for heap.leftChild(index) < heap.Size() {
		// 取出左子节点下标
		j = heap.leftChild(index)
		// 判断右孩子索引下标是否越界，如果不越界，继续判断，左右孩子哪个大。
		if j+1 < heap.Size() && heap.arr[j+1] > heap.arr[j] {
			// 此时，arr[j]将是左右孩子中最大的那个值
			j = heap.rightChild(index)
		}
		// 如果此时index 的值已经比左右两个子节点都打，说明，已经符合最大堆的条件了
		if heap.arr[index] >= heap.arr[j] {
			break
		} else {
			heap.arr[index], heap.arr[j] = heap.arr[j], heap.arr[index]
			index = j
		}
	}
}
