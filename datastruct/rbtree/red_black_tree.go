package rbtree

const (
	RED   = true
	BLACK = false
)

/*
每个节点或者是红色的，或者是黑色的；
根节点是黑色的；
每个叶子结点（最后的空节点）是黑色的；
如果一个节点是红色的，那么他的孩子都是黑色的；
从任意一个节点到叶子节点经过的黑色节点是一样的。
*/

type Node struct {
	k     int
	v     int
	left  *Node
	right *Node
	// true 代表红色， false 代表黑色
	color bool
}

func InitNode(k, v int) *Node {
	return &Node{
		k:     k,
		v:     v,
		left:  nil,
		right: nil,
		color: RED,
	}
}

type Tree struct {
	size int
	root *Node
}

func (t *Tree) Size() int {
	return t.size
}

func (t *Tree) IsEmpty() bool {
	return t.size == 0
}

func isRed(node *Node) bool {
	if node == nil {
		return BLACK
	}
	return node.color
}

//   node                     x
//  /   \     左旋转         /  \
// T1   x   --------->   node   T3
//     / \              /   \
//    T2 T3            T1   T2
func (t *Tree) leftRotate(node *Node) *Node {
	// 获取到右边的节点
	x := node.right
	// 将旋上去的节点的左子节点，赋给原先的节点
	node.right = x.left
	x.left = node
	x.color = node.color
	node.color = RED
	return x
}

// 红黑树的右旋转过程
//     node                   x
//    /   \     右旋转       /  \
//   x    T2   ------->   y   node
//  / \                       /  \
// y  T1                     T1  T2
func (t *Tree) rightRotate(node *Node) *Node {
	x := node.left
	node.left = x.right
	x.right = node
	x.color = node.color
	node.color = RED
	return x
}

// 红黑树颜色翻转
func (t *Tree) flipColors(node *Node) {
	node.color = RED
	node.left.color, node.right.color = BLACK, BLACK
}

func (t *Tree) Push(k, v int) {
	t.root = t.push(t.root, k, v)
	t.root.color = BLACK
}

func (t *Tree) push(node *Node, k, v int) *Node {
	// 因为红黑树的红节点都需要左倾，所以插入较大节点后需要进行一次左旋
	if node == nil {
		t.size++
		return InitNode(k, v)
	}
	if k > node.k {
		node.right = t.push(node.right, k, v)
	} else if k < node.k {
		node.left = t.push(node.left, k, v)
	} else {
		node.v = v
	}

	if isRed(node.right) && !isRed(node.left) {
		node = t.leftRotate(node)
	}
	if isRed(node.left) && isRed(node.left.left) {
		node = t.rightRotate(node)
	}
	if isRed(node.left) && isRed(node.right) {
		t.flipColors(node)
	}
	return node
}

func (this *Tree) getNode(node *Node, k int) *Node {
	if node == nil {
		return nil
	}
	if k == node.k {
		return node
	} else if k < node.k {
		return this.getNode(node.left, k)
	} else {
		return this.getNode(node.right, k)
	}
}
func (this *Tree) Contains(key int) bool {
	return this.getNode(this.root, key) != nil
}

func (this *Tree) GetValue(key int) *int {
	node := this.getNode(this.root, key)
	if node == nil {
		return nil
	} else {
		return &node.v
	}
}
func (this *Tree) SetNewValue(key, value int) {
	node := this.getNode(this.root, key)
	if node == nil {
		panic("没有这个key值")
	}
	node.v = value
}

func (this *Tree) minimum(node *Node) *Node {
	if node.left == nil {
		return node
	}
	return this.minimum(node.left)
}

func (this *Tree) removeMin(node *Node) *Node {
	if node.left == nil {
		rightNode := node.right
		node.right = nil
		this.size--
		return rightNode
	}
	node.left = this.removeMin(node.left)
	return node
}

func (this *Tree) Remove(node *Node, key int) *Node {
	if node == nil {
		return node
	}
	if key < node.k {
		node.left = this.Remove(node.left, key)
		return node
	} else if key > node.k {
		node.right = this.Remove(node.right, key)
		return node.right
	} else {
		if node.left == nil {
			rightNode := node.right
			node.right = nil
			this.size--
			return rightNode
		}
		if node.right == nil {
			leftNode := node.left
			node.left = nil
			this.size--
			return leftNode
		}
		successor := this.minimum(node.right)
		successor.right = this.removeMin(node.right)
		successor.left = node.left
		return successor
	}
}
