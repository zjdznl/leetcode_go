package leetcode

import "testing"

func Test_largestMultipleOfThree(t *testing.T) {
	type args struct {
		digits []int
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		// TODO: Add test cases.
		{
			name: "",
			args: args{
				digits: []int{0,0,0,0,0,0},
			},
			want: "",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := largestMultipleOfThree(tt.args.digits)
			println(got)
		})
	}
}

