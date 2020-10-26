/*
@Time : 2020/4/5 16:08
@Author : vannnnish
@File : algorithm
*/

package 二进制手表

import (
	"fmt"
	"strconv"
)

/*
Given a non-negative integer n which represents the number of LEDs that are currently on, return all possible times the watch could represent.

来源：力扣（LeetCode）
链接：https://leetcode-cn.com/problems/binary-watch
著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
*/

type ListNode struct {
	Val  int
	Next *ListNode
}

/*
	查找第
*/
func Method(num int) (res []string) {
	// count
	for i := 0; i < 12; i++ {
		for j := 0; j < 60; j++ {
			_, hourCount := convertToBin(i)
			_, minuteCount := convertToBin(j)

			if hourCount+minuteCount == num {
				if j < 10 {
					res = append(res, fmt.Sprintf("%d:0%d", i, j))
				} else {
					res = append(res, fmt.Sprintf("%d:%d", i, j))
				}
			}
		}
	}
	return res
}

// 组合公式
func turnToBinary(in string) int {
	return 0
}

// 转化十进制成二进制,并统计1的个数
func convertToBin(num int) (string, int) {
	s := ""
	count := 0
	if num == 0 {
		return "0", 0
	}
	// num /= 2 每次循环的时候 都将num除以2  再把结果赋值给 num
	for ; num > 0; num /= 2 {
		lsb := num % 2
		if lsb == 1 {
			count++
		}
		s = strconv.Itoa(lsb) + s
	}
	return s, count
}
