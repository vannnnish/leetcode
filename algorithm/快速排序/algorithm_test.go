package 快速排序

import (
	"fmt"
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
	var array = []int{4, 2, 5, 1, 7, 9}
	quickSort(array, 0, len(array)-1)
	fmt.Println(array)
}
