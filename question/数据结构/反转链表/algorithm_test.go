package 反转链表

import (
	"fmt"
	"testing"
	"time"
)

func TestReverseList(t *testing.T) {
	list := &Node{
		next: &Node{
			next: &Node{
				next: &Node{
					next: &Node{
						next: nil,
						date: 0,
					},
					date: 1,
				},
				date: 2,
			},
			date: 3,
		},
		date: 4,
	}
	newlist := reverseList(list)
	for ; newlist != nil; newlist = newlist.next {
		fmt.Println("值：", newlist.date)
		time.Sleep(1 * time.Second)
	}
}
