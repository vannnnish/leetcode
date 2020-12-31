/*
@Time : 2020/4/5 15:18
@Author : vannnnish
@File : algorithm_test
*/

package 将数之和

import (
	"testing"
)

func TestMethod(t *testing.T) {
	target := 8
	nums := []int{1, 32, 31, 14, 15, 6, 7}
	result:= Method(target, nums)
	if nums[result[0]]+nums[result[1]]==target{

	}
}
