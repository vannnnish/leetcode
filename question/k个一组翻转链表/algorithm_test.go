package k个一组翻转链表

import (
	"fmt"
	"testing"
)

func TestMethod(t *testing.T) {
	var list = &ListNode{
		Val: 0,
		Next: &ListNode{
			1,
			&ListNode{
				Val: 2,
				Next: &ListNode{
					Val: 3,
					Next: &ListNode{
						Val: 4,
						Next: &ListNode{
							Val: 5,
							Next: &ListNode{
								Val:  6,
								Next: nil,
							},
						},
					},
				},
			},
		},
	}

	res:=Method(list,2)
	fmt.Println("res:",res)
	if res!=nil{
		for ; res != nil; {
			fmt.Println(res.Val)
			res = res.Next
		}
	}
}
