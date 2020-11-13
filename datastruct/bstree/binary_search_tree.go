package bstree

import (
	"fmt"
	"strings"
)

type Node struct {
	E     int
	Left  *Node
	Right *Node
}

// 约定在此样例代码中我们的二分搜索树中没有重复元素
// 如果想包涵重复元素的话，只需要以下定义：
// 左子树小于等于此节点，或右子树大于等于节点
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

func InitNode(E int) *Node {
	return &Node{
		E:     E,
		Left:  nil,
		Right: nil,
	}
}

func (t *Tree) AddE(e int) {
	t.root = t.add(t.root, e)
}

// 向以node为根的二分搜索树中插入元素E，递归算法
func (t *Tree) add(node *Node, e int) *Node {

	// 不管是递归还是回溯，首先我们都应该先写出递归的结束条件是什么
	if node == nil {
		t.size++
		return InitNode(e)
	}

	if e > node.E {
		node.Right = t.add(node.Right, e)
	} else if e < node.E {
		node.Left = t.add(node.Left, e)
	}
	return node
}

// 查找二分搜索中是否含有元素E
func (t *Tree) Contains(e int) bool {
	return t.contains(t.root, e)
}

// 递归的方式查找元素是否存在
func (t *Tree) contains(node *Node, e int) bool {
	if node == nil {
		return false
	}
	if e == node.E {
		return true
	} else if e > node.E {
		return t.contains(node.Right, e)
	} else {
		return t.contains(node.Left, e)
	}
}

// 遍历算法
// 1.前序遍历
func (t *Tree) PreOrder() {
	PreOrder(t.root)
	fmt.Println()
}
func PreOrder(node *Node) {
	if node == nil {
		return
	}
	fmt.Printf("%d ", node.E)
	PreOrder(node.Left)
	PreOrder(node.Right)
}

// 非递归的前序遍历
func (t *Tree) PreOrderNR() {
	stack := make([]*Node, 0)
	stack = append(stack, t.root)
	for len(stack) > 0 {
		curNode := stack[len(stack)-1]
		stack = stack[:len(stack)-1]
		fmt.Printf("%d ", curNode.E)
		if curNode.Right != nil {
			stack = append(stack, curNode.Right)
		}
		if curNode.Left != nil {
			stack = append(stack, curNode.Left)
		}
	}
	fmt.Println()
}

// 2.中序遍历
func (t *Tree) MidOrder() {
	MidOrder(t.root)
}
func MidOrder(node *Node) {
	if node == nil {
		return
	}

	MidOrder(node.Left)
	fmt.Printf("%d ", node.E)
	MidOrder(node.Right)
}

// 3.后序遍历
func (t *Tree) BackOrder() {
	BackOrder(t.root)
}
func BackOrder(node *Node) {
	if node == nil {
		return
	}
	BackOrder(node.Left)
	BackOrder(node.Right)
	fmt.Printf("%d ", node.E)
}

// 二分搜索树的层序遍历
func (t *Tree) LevelOrder() {
	queue := make([]*Node, 0)
	queue = append(queue, t.root)
	for len(queue) > 0 {
		curNode := queue[0]
		queue = queue[1:]
		fmt.Printf("%d ", curNode.E)
		if curNode.Left != nil {
			queue = append(queue, curNode.Left)
		}
		if curNode.Right != nil {
			queue = append(queue, curNode.Right)
		}
	}
}

// 二分搜索树中搜索最小值
func (t *Tree) FindMin() int {
	if t.IsEmpty() {
		panic("二叉树为空，无法删除任何节点")
	}
	return minimum(t.root).E
}
func minimum(node *Node) *Node {
	if node.Left == nil {
		return node
	}
	return minimum(node.Left)
}

// 二分搜索树中搜索最大值
func (t *Tree) FindMax() int {
	if t.IsEmpty() {
		panic("二叉树为空，无法删除任何节点")
	}
	return maximum(t.root).E
}
func maximum(node *Node) *Node {
	if node.Right == nil {
		return node
	}
	return maximum(node.Right)
}

// 从二分搜索树中删除最小值
func (t *Tree) DelMin() int {
	var ret int = t.FindMin()
	t.root = t.rmMin(t.root)
	return ret
}

// 删除掉以node为根的二分搜索树的最小节点
// 返回删除节点后新的二分搜索树的根
func (t *Tree) rmMin(node *Node) *Node {
	if node.Left == nil {
		nodeRight := node.Right
		node.Right = nil
		t.size--
		return nodeRight
	}
	node.Left = t.rmMin(node.Left)
	return node
}

// 从二分搜索树种删除最大值
func (t *Tree) DelMax() int {
	var ret int = t.FindMax()
	t.root = t.rmMax(t.root)
	return ret
}

// 删除掉以node为根的二分搜索树的最小节点
// 返回删除节点后新的二分搜索树的根
func (t *Tree) rmMax(node *Node) *Node {
	if node.Right == nil {
		nodeLeft := node.Left
		node.Left = nil
		t.size--
		return nodeLeft
	}
	node.Right = t.rmMax(node.Right)
	return node
}

// 在二分搜索树中删除值为e的方法
func (t *Tree) Remove(e int) {
	t.root = t.remove(t.root, e)
}
func (t *Tree) remove(node *Node, e int) *Node {
	if node == nil {
		return nil
	}
	if e > node.E {
		node.Right = t.remove(node.Right, e)
		return node
	} else if e < node.E {
		node.Left = t.remove(node.Left, e)
		return node
	} else {
		// 如果左子树为空的时候
		if node.Left == nil {
			nodeRight := node.Right
			node.Right = nil
			t.size--
			return nodeRight
		}
		// 如果右子树为空
		if node.Right == nil {
			nodeLeft := node.Left
			node.Left = nil
			t.size--
			return nodeLeft
		}
		// 如果左右子树都不为空，那么我们需要找到node的后继
		// 就是所有比node值大的节点中值最小的那个，显然它在node的右子树中
		nodeNext := minimum(node.Right)
		nodeNext.Right = t.rmMin(node.Right)
		nodeNext.Left = node.Left
		node.Left = nil
		node.Right = nil
		return nodeNext
	}
}

// 重构二叉树的打印方法
func (t *Tree) String() string {
	var (
		builder strings.Builder
	)
	generateBSTString(t.root, 0, &builder)
	return builder.String()
}
func generateBSTString(node *Node, depth int, builder *strings.Builder) {
	if node == nil {
		fmt.Fprintln(builder, generateDepthString(depth)+"null")
		return
	}
	fmt.Fprintln(builder, generateDepthString(depth), node.E)
	generateBSTString(node.Left, depth+1, builder)
	generateBSTString(node.Right, depth+1, builder)
}
func generateDepthString(depth int) string {
	var builder strings.Builder
	for i := 0; i < depth; i++ {
		fmt.Fprintf(&builder, "--")
	}
	return builder.String()
}
