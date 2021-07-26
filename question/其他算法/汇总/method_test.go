package 汇总

import (
	"fmt"
	"testing"
)

func Test_kthSmallest(t *testing.T) {
	input := [][]int{{1, 3, 11}, {2, 4, 6}}
	fmt.Println(kthSmallest(input, 5))

}

func Test_IsValidBST(t *testing.T) {
	v := &TreeNode{
		Val: 5,
		Left: &TreeNode{
			Val:   1,
			Left:  nil,
			Right: nil,
		},
		Right: &TreeNode{
			Val: 4,
			Left: &TreeNode{
				Val:   3,
				Left:  nil,
				Right: nil,
			},
			Right: &TreeNode{
				Val:   6,
				Left:  nil,
				Right: nil,
			},
		},
	}
	/*	v := &TreeNode{
		Val: 2,
		Left: &TreeNode{
			Val:   1,
			Left:  nil,
			Right: nil,
		},
		Right: &TreeNode{
			Val:   3,
			Left:  nil,
			Right: nil,
		},
	}*/
	fmt.Println(isValidBST(v))
}

func Test_IsValidBST1(t *testing.T) {
	v := &TreeNode{
		Val:  2,
		Left: nil,
		Right: &TreeNode{
			Val:   2,
			Left:  nil,
			Right: nil,
		},
	}
	fmt.Println(isValidBST(v))
}

func Test_findLastestStep(t *testing.T) {
	input := []int{3, 5, 1, 2, 4}
	m := 1
	findLatestStep(input, m)
}

