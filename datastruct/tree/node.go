package tree

// 节点的数据结构
type Node struct {
	E     int
	Left  *Node
	Right *Node
}

// 树的数据结构，是不是和链表的定义十分的相似
type Tree struct {
	root *Node
	size int
}

func (t Tree) Size() int {
	return t.size
}

func (t Tree) Root() *Node {
	return t.root
}

// 判断二叉树是否为空
func (t *Tree) IsEmpty() bool {
	if t.size == 0 {
		return true
	}
	return false
}

// 节点的初始化方法
func InitNode(E int) *Node {
	return &Node{
		E:     E,
		Left:  nil,
		Right: nil,
	}
}
