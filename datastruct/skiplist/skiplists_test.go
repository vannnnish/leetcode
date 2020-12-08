package skiplist

import (
	"fmt"
	"testing"
)

func Test_randomLevel(t *testing.T) {
	var valueMap = map[int]int{}
	for i := 0; i < 1000000; i++ {
		v := randomLevel()
		valueMap[v] = valueMap[v] + 1
	}
	for k, v := range valueMap {
		fmt.Printf("v:%d,count:%d", k, v)
		fmt.Println()
	}
}

func TestSkipList(t *testing.T) {
	sl := New()

	sl.Insert(float64(100), "foo")

	e, ok := sl.Search(float64(100))
	// 打印是否能从跳表找到100，应该输出 true
	fmt.Println(ok)
	// 100 对应的值，这里应为 foo
	fmt.Println(e.Value)
	e, ok = sl.Search(float64(200))
	// 打印是否能从跳表找到200，应该输出 false
	fmt.Println(ok)
	// 200 对应的值，这里应为 nil
	fmt.Println(e)

	sl.Insert(float64(20.5), "bar")
	sl.Insert(float64(50), "spam")
	sl.Insert(float64(20), 42)
	// 打印跳表长度，应输出4
	fmt.Println(sl.Size())
	e = sl.Delete(float64(50))
	// 50 对应的值，应该为 spam
	fmt.Println(e.Value)
	// 打印跳表长度，应输出 3
	fmt.Println(sl.Size())
	// 打印跳表所有元素的 Value
	for e := sl.Front(); e != nil; e = e.Next() {
		fmt.Println(e.Value)
	}
}
