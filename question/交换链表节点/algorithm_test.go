/*
@Time : 2020/4/5 15:18
@Author : vannnnish
@File : algorithm_test
*/

package 交换链表节点

import (
	"fmt"
	"testing"
)

func TestMethod(t *testing.T) {
	var list = &ListNode{
		Val: 0,
		Next: &ListNode{
			Val: 1,
			Next: &ListNode{
				Val: 2,
				Next: &ListNode{
					Val: 3,
					Next: &ListNode{
						Val:  4,
						Next: nil,
					},
				},
			},
		},
	}

	result := Method(list)
	for {
		if result != nil {
			fmt.Println(result.Val)
			result = result.Next
		} else {
			break
		}
	}
}
