package leetcode

import "testing"

func Test_numTimesAllBlue(t *testing.T) {
	type args struct {
		light []int
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
				light: []int{2,1,3,5,4},
			},
			want: 0,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := numTimesAllBlue(tt.args.light); got != tt.want {
				t.Errorf("numTimesAllBlue() = %v, want %v", got, tt.want)
			}
		})
	}
}
