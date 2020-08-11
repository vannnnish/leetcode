package 堆排序

/*
	parent=(i-1)/2
	child1=2i+1
	child2=2i+2
*/

// 构造堆过程
func heapify(tree []int, n, root int) {
	child1 := root*2 + 1
	child2 := root*2 + 2
	var max = root
	if child1 < n && tree[child1] > tree[max] {
		max = child1
	}
	if child2 < n && tree[child2] > tree[max] {
		max = child2
	}
	if max != root {
		tree[root], tree[max] = tree[max], tree[root]
		heapify(tree, n, max)
	}

}
