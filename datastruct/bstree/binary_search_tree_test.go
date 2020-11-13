package bstree

import (
	"fmt"
	"math/rand"
	"testing"
	"time"
)

func TestBSTreeTraverse(t *testing.T) {
	m := 100000000
	tree := &Tree{}
	startTime := time.Now()
	source := rand.NewSource(time.Now().Unix())
	r := rand.New(source)
	for i := 0; i < m; i++ {
		tree.AddE(r.Intn(m))
	}
	//tree.PreOrder()
	//tree.PreOrderNR()
	insertTime := time.Now()
	for i := 0; i < m; i++ {
		tree.Contains(r.Intn(m))
	}
	endTime := time.Now()
	fmt.Println("对于千万级数据，二叉搜索树插入共用时：", insertTime.Sub(startTime))
	fmt.Println("对于千万级数据，二叉搜索树查找共用时：", endTime.Sub(insertTime))
}
