package rbtree

import (
	"fmt"
	"math"
	"math/rand"
	"testing"
	"time"
)

func TestRedBlackTree(t *testing.T) {
	m := 1000000
	rbtree := &Tree{
		root: nil,
		size: 0,
	}
	source := rand.NewSource(time.Now().Unix())
	r := rand.New(source)
	startTime := time.Now()
	for i := 0; i < m; i++ {
		rbtree.Push(r.Intn(m), rand.Intn(math.MaxInt32))
	}
	insertTime := time.Now()
	for i := 0; i < m; i++ {
		rbtree.Contains(r.Intn(m))
	}
	endTime := time.Now()
	fmt.Println("对于千万级数据，红黑树插入共用时：", insertTime.Sub(startTime))
	fmt.Println("对于千万级数据，红黑树查找共用时：", endTime.Sub(insertTime))
}
