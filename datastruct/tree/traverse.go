package tree

import "fmt"

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

func (t *Tree) InOrder() {
	InOrder(t.root)
	fmt.Println()
}

// 中序遍历
func InOrder(node *Node) {
	if node == nil {
		return
	}
	InOrder(node.Left)
	fmt.Printf("%d ", node.E)
	InOrder(node.Right)
}

// 后续遍历
func (t *Tree) PostOrder() {
	PostOrder(t.root)
	fmt.Println()
}

func PostOrder(node *Node) {
	if node == nil {
		return
	}
	PostOrder(node.Left)
	PostOrder(node.Right)
	fmt.Printf("%d ", node.E)
}

// 广度优先，层序遍历
func (t *Tree) LevelOrder() {
	OutputNode(t.root)
	fmt.Println()
}

// 输出节点
func OutputNode(node *Node) {
	if node == nil {
		return
	}
	var nodeList = make([]*Node, 0)
	nodeList = []*Node{node}
	var index int
	for ; ; {
		length := len(nodeList)
		var hasChild = false
		for i := index; i < length; i++ {
			if nodeList[i] != nil {
				hasChild = true
				nodeList = append(nodeList, nodeList[i].Left)
			}
			if nodeList[i] != nil {
				hasChild = true
				nodeList = append(nodeList, nodeList[i].Right)
			}
			index++
		}
		if !hasChild {
			break
		}

	}
	for i := 0; i < len(nodeList); i++ {
		if nodeList[i] == nil {
			continue
		}
		fmt.Printf("%d ", nodeList[i].E)
	}
}

// 树的非递归遍历
func (t *Tree) PreOrderNR() {
	stack := make([]*Node, 0)
	stack = append(stack, t.root)
	for len(stack) > 0 {
		// 获取到当前的节点
		curNode := stack[len(stack)-1]
		// 将这个节点出栈, 此时
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
