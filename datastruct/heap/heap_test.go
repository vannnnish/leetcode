package heap

import (
	"fmt"
	"math"
	"math/rand"
	"testing"
	"time"
)

func TestHeap(t *testing.T) {
	n := 1000000
	ArrHeap := MaxHeap{
		arr: make([]int, 0),
	}
	start := time.Now()
	source := rand.NewSource(start.Unix())
	r := rand.New(source)
	for i := 0; i < n; i++ {
		v := r.Intn(math.MaxInt32)
		ArrHeap.Add(v)
	}
	// 判断生成的堆是否合法
	for i := 0; i < n; i++ {
		if 2*i+2 < ArrHeap.Size() {
			break
		}
		if ArrHeap.arr[i] < ArrHeap.arr[2*i+1] || ArrHeap.arr[i] < ArrHeap.arr[2*i+2] {
			panic("生成的堆不合法")
		}
	}
	// 生成一个一百万空间的数组把堆装回去，如果我们的算法无误，那么数组将是有序的
	test_arr := make([]int, n)
	for i := 0; i < n; i++ {
		test_arr[i] = ArrHeap.RemoveMax()
	}
	end := time.Now()
	for i := 1; i < n; i++ {
		if test_arr[i-1] < test_arr[i] {
			panic(fmt.Sprintf("第%d个元素开始出现无序", i))
		}
	}
	fmt.Printf("大顶堆运行成功！整理%d万级数据共用时：%dms", n/10000, end.Sub(start).Milliseconds())
}

func TestArr(t *testing.T) {
	originArr := []int{1, 2, 3, 4, 5, 6, 7}
	halfArr := originArr[:len(originArr)-1]
	fmt.Println(halfArr)
}
