package 快速排序

import (
	"fmt"
	"leetcode/algorithm/common"
	"testing"
)

func TestMethod(t *testing.T) {
	type args struct {
		array []int
		begin int
		end   int
		mark  string
	}
	tests := []struct {
		name string
		args args
	}{
		{name: "xixi", args: args{
			array: []int{4, 2, 5, 1, 7, 9},
			begin: 0,
			end:   5,
			mark:  "--",
		}},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			Method(tt.args.array, tt.args.begin, tt.args.end, tt.args.mark)
			fmt.Println(tt.args.array)
		})
	}
}

func TestQuickSort(t *testing.T) {
	var array = []int{4, 2, 5, 2, 3, 4, 1, 7, 45, 7}
	quickSort4(array)
	fmt.Println(array)

}

func TestShellSortFromOnline(t *testing.T) {
	var array = common.RandomArrGenerator(100000)
	common.DoTimeStatic(array, quickSort3)
	fmt.Println(common.IsValidSort(array))
}

func TestQuickSort4(t *testing.T) {
	//var array = common.RandomArrGenerator(10000000)
	//var array = common.InOrderArrGenerator(1000000)
	var array = common.RepeatNumberArrGenerator(1000000, 10)
	common.DoTimeStatic(array, quickSort3)
	common.DoTimeStatic(array, quickSort4)
	fmt.Println(common.IsValidSort(array))
}

func TestQuickSort5(t *testing.T) {
	var array1 = []int{4, 2, 5, 2, 3, 4, 1, 7, 45, 7}
	quickSort5(array1)
	fmt.Println(array1)
	//var array = common.RandomArrGenerator(10000000)
	//var array = common.InOrderArrGenerator(1000000)
	var array = common.RepeatNumberArrGenerator(10000000, 10)
	common.DoTimeStatic(array, quickSort5)
	fmt.Println(common.IsValidSort(array))
}
