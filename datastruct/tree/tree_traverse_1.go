package tree

import "fmt"

type Stack struct {
	V [15]*Node
}

func (s *Stack) IsEmpty() bool {
	if s.V[0] == nil {
		return true
	}
	return false
}

func (s *Stack) Push(n *Node) bool {
	// 寻找插入点
	insertIndex := -1
	for i := range s.V {
		if s.V[i] == nil {
			insertIndex = i
			break
		}
	}
	if insertIndex == -1 {
		return false
	}
	s.V[insertIndex] = n
	return true
}

func (s *Stack) Pop() (*Node, bool) {
	for i := range s.V {
		if i == 0 && s.V[0] == nil {
			return nil, false
		}
		if s.V[i] == nil && i != 0 {
			return s.V[i-1], true
		}
	}
	return s.V[14], true
}

// 前序遍历
func (t *Tree) PreOrder1() (res []int) {
	s := &Stack{}
	node := t.root
	for node != nil || !s.IsEmpty() {
		for node != nil {
			res = append(res, node.E)
			s.Push(node)
			node = node.Left
		}
		node, _ = s.Pop()
		node = node.Right
	}
	fmt.Println(res)
	return res
}

type TreeNode struct {
	value       int
	left, right *TreeNode
}

func inorder(root *TreeNode) (res []int) {
	var stack []*TreeNode
	for root != nil || len(stack) > 0 {
		for root != nil {
			stack = append(stack, root)
			root = root.left
		}
		root = stack[len(stack)-1]
		stack = stack[:len(stack)-1]
		res = append(res, root.value)
		root = root.right
	}
	return res
}

func preOrder(root *TreeNode) (res []int) {
	var stack []*TreeNode
	for root != nil || len(stack) > 0 {
		for root != nil {
			res = append(res, root.value)
			stack = append(stack, root)
			root = root.left
		}
		root = stack[len(stack)-1].right
		stack = stack[:len(stack)-1]
	}
	return
}

func main() {
	var node TreeNode
	node = TreeNode{1, nil, nil}
	node.left = &TreeNode{2, nil, nil}
	node.right = &TreeNode{3, nil, nil}
	fmt.Println(inorder(&node))
}
