/*
@Time : 2020/4/5 15:18
@Author : vannnnish
@File : algorithm
*/

package 交换链表节点

type ListNode struct {
	Val  int
	Next *ListNode
}

func Method(head *ListNode) *ListNode {
	dummyHead := &ListNode{0, head}
	temp := dummyHead
	for temp.Next != nil && temp.Next.Next != nil {
		node1 := temp.Next
		node2 := temp.Next.Next
		temp.Next = node2
		node1.Next = node2.Next
		node2.Next = node1
		temp = node1
	}
	return dummyHead.Next
}
