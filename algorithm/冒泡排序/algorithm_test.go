package 选择排序

import (
	"fmt"
	"leetcode/algorithm/common"
	"testing"
)

func TestBubbleSort(t *testing.T) {
	var array = common.RandomArrGenerator(50000)
	common.DoTimeStatic(array, bubbleSort)
	fmt.Println(common.IsValidSort(array))
}
