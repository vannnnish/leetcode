package 汇总

import "math"

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func isValidBST(root *TreeNode) bool {
	a := true
	if root == nil  {
		return true
	}
	preRecrive(root, true, math.MaxInt64, &a)
	return a
}

// 中序遍历
func preRecrive(root *TreeNode, isLeft bool, parentVal int, flag *bool) {
	if root == nil {
		return
	}
	preRecrive(root.Left, true, root.Val, flag)
	if isLeft {
		if root.Val >= parentVal {
			*flag = false
			return
		}
	} else {
		if root.Val <= parentVal {
			*flag = false
			return
		}
	}
	preRecrive(root.Right, false, root.Val, flag)
	return
}
