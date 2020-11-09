package tree

import "testing"

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
