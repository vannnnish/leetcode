package queue

import (
	"fmt"
	"testing"
)

func TestQueue_Queue(t *testing.T) {
	queue := NewQueue(10)

	for i := 0; i < 12; i++ {
		fmt.Println(queue.EnQueue(i))
	}

	for i := 0; i < 12; i++ {
		fmt.Println(queue.DeQueue())
	}
}
