/*
@Time : 2020/4/5 15:18
@Author : vannnnish
@File : algorithm
*/

package 字符串转整数

import (
	"math"
	"strings"
)

func Method(s string) int {
	str := strings.TrimSpace(s)
	var result = 0
	var sign = 0
	for i, v := range str {
		if v >= '0' && v <= '9' {
			result = result*10 + int(v-'0')
		} else if v == '-' && i == 0 {
			sign = -1
		} else if v == '+' && i == 0 {
			sign = 1
		} else {
			break
		}

		if result > math.MaxInt32 {
			if sign == -1 {
				return math.MinInt32
			}
			return math.MaxInt32
		}
	}
	return result * sign
}
