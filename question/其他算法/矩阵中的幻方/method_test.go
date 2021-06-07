package 矩阵中的幻方

import (
	"fmt"
	"testing"
)

func Test_numMagicSquaresInside(t *testing.T) {
	var arr = [][]int{
		{4, 3, 8, 4},
		{9, 5, 1, 9},
		{2, 7, 6, 2},
	}
	fmt.Println(numMagicSquaresInside(arr))
}

func Test_toArr(t *testing.T) {
	var arr []int
	arr = []int{3, 30, 34,  5, 9}
	toSortedStr(arr)
}
