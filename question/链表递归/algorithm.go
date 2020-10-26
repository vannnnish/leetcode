/*
@Time : 2020/4/5 16:08
@Author : vannnnish
@File : algorithm
*/

package 链表递归

import "fmt"

type ListNode struct {
	Val  int
	Next *ListNode
}

/*
 	查找第
*/
func Method(l1 *ListNode, k int) (final int) {
	//var final = 0
	ok:=false
	defer func(val *int) {
		if err := recover(); err != nil {
			fmt.Println("捕获异常:", err)
			*val,ok=err.(int)
			fmt.Println("ok:",ok)
			fmt.Println(*val)
		}
	}(&final)
	recursion(l1.Next, k)
	return final
}

var i = 1

func recursion(current *ListNode, k int) {
	if current == nil {
		return
	}
	currentVal := current.Val
	recursion(current.Next, k)
	if i == k {
		panic(currentVal)
	}
	i++
}
