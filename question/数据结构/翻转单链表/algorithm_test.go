/*
@Time : 2020/4/5 16:08
@Author : vannnnish
@File : algorithm_test
*/

package 翻转单链表

import "testing"

func TestRotate(t *testing.T) {
	rotateRight(&ListNode{
		1, &ListNode{
			2,
			&ListNode{
				Val:  0,
				Next: nil,
			},
		},
	},1)
}
