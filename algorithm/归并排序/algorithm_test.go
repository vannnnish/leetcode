package 归并排序

import (
	"fmt"
	"leetcode/algorithm/common"
	"testing"
)

func TestMergeSort(t *testing.T) {
	var array = common.RandomArrGenerator(5000000)
	common.DoTimeStatic(array, mergeSortFromOnline)
	array = mergeSort(array)
	fmt.Println(common.IsValidSort(array))
}

func TestMerge(t *testing.T) {
	var array = common.InOrderArrGenerator(5000000)
	//var array = common.RandomArrGenerator(5000000)
	common.DoTimeStatic(array, mergeSortByMyself)
	array = apart(array)
	fmt.Println(common.IsValidSort(array))
}
