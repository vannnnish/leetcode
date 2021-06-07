package 选择排序

import (
	"fmt"
	"leetcode/algorithm/common"
	"testing"
)

func TestChooseSort(t *testing.T) {
	var array = common.RandomArrGenerator(50000)
	choseSort(array)
	common.DoTimeStatic(array,choseSort)
	fmt.Println(common.IsValidSort(array))
}
