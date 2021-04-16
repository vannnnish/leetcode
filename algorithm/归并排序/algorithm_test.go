package 归并排序

import (
	"fmt"
	"leetcode/algorithm/common"
	"testing"
)

func TestMergeSort(t *testing.T) {
	var array = common.RandomArrGenerator(500000)
	common.DoTimeStatic(array, mergeSortFromOnline)
	array = mergeSort(array)
	fmt.Println(common.IsValidSort(array))
}
