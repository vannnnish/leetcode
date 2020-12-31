/*
@Time : 2020/4/5 16:08
@Author : vannnnish
@File : algorithm
*/

package 两数相加

type ListNode struct {
	Val  int
	Next *ListNode
}

func Method(l1 *ListNode, l2 *ListNode) *ListNode {
	return addNode(l1, l2, 0)
}

func addNode(l1 *ListNode, l2 *ListNode, pre int) *ListNode {
	if l1 == nil && l2 == nil {
		if pre != 0 {
			return &ListNode{Val: pre}
		} else {
			return nil
		}
	}

	if l1 == nil {
		l1 = &ListNode{Val: 0}
	}
	if l2 == nil {
		l2 = &ListNode{Val: 0}
	}
	totalSum := l1.Val + l2.Val + pre
	i, j := totalSum/10, totalSum%10
	l := new(ListNode)
	l.Val = j
	l.Next = addNode(l1.Next, l2.Next, i)
	return l
}
