package leetcode

import "testing"

func Test_sumFourDivisors(t *testing.T) {
	type args struct {
		nums []int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		// TODO: Add test cases.
		{
			name: "",
			args: args{
				nums: []int{1,2,3,4,5,6,7,8,9,10},
			},
			want: 0,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := sumFourDivisors(tt.args.nums); got != tt.want {
				t.Errorf("sumFourDivisors() = %v, want %v", got, tt.want)
			}
		})
	}
}
