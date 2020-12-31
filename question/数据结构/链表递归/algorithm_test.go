/*
@Time : 2020/4/5 16:08
@Author : vannnnish
@File : algorithm_test
*/

package 链表递归

import (
	"fmt"
	"testing"
)

func TestMethod(t *testing.T) {
	l1:=&ListNode{
		Val:  1,
		Next: &ListNode{
			Val: 2,
			Next: &ListNode{
				Val: 3,
				Next: &ListNode{
					Val:  4,
					Next: &ListNode{
						Val: 7,
						Next: nil,
					},
				},
			},
		},
	}
	fmt.Println(Method(l1,3))
}