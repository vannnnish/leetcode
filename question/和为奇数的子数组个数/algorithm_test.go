package 和为奇数的子数组个数

import (
	"fmt"
	"testing"
)

func Test_combination(t *testing.T) {
	fmt.Println(combination(10, 2))
}

func Test_numOfSubarrays(t *testing.T) {
	arr := []int{1, 2, 3, 4, 5, 6, 7}
	fmt.Println(numOfSubarrays(arr))
}
