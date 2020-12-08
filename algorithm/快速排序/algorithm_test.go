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
			array: []int{1, 2, 4, 65, 7, 8, 9, 23},
			begin: 0,
			end:   2,
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
