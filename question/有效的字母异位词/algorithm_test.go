package 有效的字母异位词

import "testing"

func TestMethod(t *testing.T) {
	type args struct {
		str1 string
		str2 string
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		// TODO: Add test cases.
		{
			"测试1",
			args{
				"abcde",
				"bcdae",
			},
			true,
		},
		{
			"测试2",
			args{
				"abcde2",
				"bcdae",
			},
			true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Method(tt.args.str1, tt.args.str2); got != tt.want {
				t.Errorf("Method() = %v, want %v", got, tt.want)
			}
		})
	}
}
