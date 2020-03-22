package leetcode

import "testing"

func Test_hasValidPath(t *testing.T) {
	type args struct {
		grid [][]int
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		// TODO: Add test cases.
		{
			name: "",
			args: args{
				grid: [][]int{
					{1,1,2},
				},
			},
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := hasValidPath(tt.args.grid); got != tt.want {
				t.Errorf("hasValidPath() = %v, want %v", got, tt.want)
			}
		})
	}
}
