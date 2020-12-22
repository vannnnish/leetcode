/*
@Time : 2020/4/5 16:08
@Author : vannnnish
@File : algorithm
*/

package 翻转单链表

import (
	"fmt"
)


type ListNode struct {
	Val  int
	Next *ListNode
}

func rotateRight(head *ListNode, k int) *ListNode {
	start := head
	originStart := head
	// 获取链表长度
	length := 0
	for ; head != nil; head = head.Next {
		length++
	}
	if length <= 1 {
		return originStart
	}
	mod := k % length
	if mod == 0 {
		return originStart
	}
	//var cutListNode *ListNode
	// 找到截断点
	fmt.Println(start)
	var cut *ListNode
	var originCut *ListNode
	for i := 0; i < length-mod; i++ {
		if i == length-mod-1 {
			cut = start.Next
			start.Next = nil
		} else {
			start = start.Next
		}
	}
	originCut = cut
	for ; cut.Next != nil; cut = cut.Next {
	}
	cut.Next = originStart
	return originCut
}
