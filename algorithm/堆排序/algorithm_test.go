package 堆排序

import (
	"fmt"
	"math"
	"testing"
)

func Test_heapify(t *testing.T) {
	heap := NewHeap(5)
	heap.Insert(1)
	heap.Insert(2)
	heap.Insert(9)
	heap.Insert(3)
	heap.Insert(4)
	heap.Insert(7)
	heap.Insert(5)
	fmt.Println(heap.data)
	fmt.Println(heap.Remove())
	fmt.Println(heap.Remove())
	fmt.Println(heap.Remove())
	fmt.Println(heap.Remove())
	fmt.Println(heap.Remove())
	fmt.Println(heap.Remove())
	fmt.Println(heap.Remove())
}

func TestName(t *testing.T) {
	fmt.Println(isSqrt(1))
	fmt.Println(isSqrt(2))
	fmt.Println(isSqrt(3))
	fmt.Println(isSqrt(4))
}

func isSqrt(n int) bool {
	sqrt := math.Sqrt(float64(n))
	str:="123"
	fmt.Println(str[0:len(str)])
	return n == int(sqrt)*int(sqrt)
}
