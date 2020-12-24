package 反转链表

type Node struct {
	next *Node
	date int
}

func reverseList(head *Node) *Node {
	if head == nil || head.next == nil {
		return head
	}
	list := reverseList(head.next)
	head.next.next = head
	head.next = nil
	return list
}
