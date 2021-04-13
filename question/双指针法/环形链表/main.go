package 环形链表

type ListNode struct {
	Val  int
	Next *ListNode
}

func detectCycle(head *ListNode) *ListNode {
	slow, fast := head, head
	//  如果快指针不为空
	for fast != nil {
		slow = slow.Next
		// 如果fast指针下一个节点为空，表明无环
		if fast.Next == nil {
			return nil
		}
		// 快指针遍历
		fast = fast.Next.Next
		// 如果相等这，遇到了环
		if fast == slow {
			// 将慢指针移动到链表头部
			p := head
			for p != slow {
				p = p.Next
				slow = slow.Next
			}
			return p
		}
	}
	return nil
}
