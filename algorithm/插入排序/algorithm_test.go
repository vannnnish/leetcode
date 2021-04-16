package 插入排序

import (
	"fmt"
	"leetcode/algorithm/common"
	"testing"
)

func TestChooseSort(t *testing.T) {
	var array = common.RandomArrGenerator(50000)
	common.DoTimeStatic(array, insertSort)
	fmt.Println(common.IsValidSort(array))
}

func TestChooseSortOptimization(t *testing.T) {
	var array = common.RandomArrGenerator(100000)
	common.DoTimeStatic(array, insertSortOptimization)
	fmt.Println(common.IsValidSort(array))
}

func TestShellSort(t *testing.T) {
	var array = common.RandomArrGenerator(100000)
	common.DoTimeStatic(array, shellSort)
	fmt.Println(common.IsValidSort(array))
}

func TestShellSortFromOnline(t *testing.T) {
	var array = common.RandomArrGenerator(10000000)
	common.DoTimeStatic(array, ShellSort)
	fmt.Println(common.IsValidSort(array))
}
