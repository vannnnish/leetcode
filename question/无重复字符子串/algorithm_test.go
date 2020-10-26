package 无重复字符子串

import (
	"testing"
)

func TestMethod(t *testing.T) {
	type args struct {
		s string
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		// TODO: Add test cases.
		{name: "test1", args: args{s:"abcabcbb"}, want:3 },
		{name: "test1", args: args{s:"ab"}, want:2 },
		{name: "test1", args: args{s:"aab"}, want:2 },
		{name: "test2", args: args{s:""}, want:0 },
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Method(tt.args.s); got != tt.want {
				t.Errorf("Method() = %v, want %v", got, tt.want)
			}
		})
	}
}