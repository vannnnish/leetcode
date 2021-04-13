package tree

import (
	"fmt"
	"testing"
)

// 这里有打印操作，所以别忘了import "fmt"
func TestTree(t *testing.T) {
	tree := &Tree{}
	tree.root = InitNode(0)
	tree.root.Left = InitNode(3)
	tree.root.Right = InitNode(4)
	tree.root.Left.Left = InitNode(1)
	tree.root.Left.Right = InitNode(6)
	tree.root.Right.Left = InitNode(2)
	tree.root.Right.Right = InitNode(9)
	tree.root.Right.Right.Right = InitNode(10)

	tree.root.Left.Right.Left = InitNode(7)
	tree.root.Left.Right.Right = InitNode(5)
	tree.PreOrder()
	tree.InOrder()
	tree.PostOrder()
	tree.LevelOrder()
}

// 这里有打印操作，所以别忘了import "fmt"
func TestTree1(t *testing.T) {
	tree := &Tree{}
	tree.root = InitNode(0)
	tree.root.Left = InitNode(3)
	tree.root.Right = InitNode(4)
	tree.root.Left.Left = InitNode(1)
	tree.root.Left.Right = InitNode(6)
	tree.root.Right.Left = InitNode(2)
	tree.root.Right.Right = InitNode(9)
	tree.root.Right.Right.Right = InitNode(10)

	tree.root.Left.Right.Left = InitNode(7)
	tree.root.Left.Right.Right = InitNode(5)
	//tree.PreOrder()
	fmt.Println()
	tree.PreOrder1()

}

func TestInorder(t *testing.T) {
	var node TreeNode
	node = TreeNode{1, nil, nil}
	node.left = &TreeNode{2, &TreeNode{
		value: 8,
		left:  nil,
		right: nil,
	}, &TreeNode{
		value: 3,
	}}
	node.right = &TreeNode{3, &TreeNode{
		value: -1,
		left:  nil,
		right: nil,
	}, nil}
	fmt.Println(preOrder(&node))
}

func TestIsPow(t *testing.T) {
	fmt.Println(isPowOf(8, 2))
	fmt.Println(isPowOf(1, 2))
	fmt.Println(isPowOf(2, 2))
	fmt.Println(isPowOf(9, 3))
	fmt.Println(isPowOf(8, 3))
	fmt.Println(isPowOf(10, 2))
	fmt.Println(isPowOf(27, 3))
}

func isPowOf(v int, pow int) bool {

	for {
		if v == pow {
			return true
		} else if v == 1 {
			return false
		}
		if v%pow != 0 {
			return false
		}
		v = v / pow
	}
}
