package leetcode

import "testing"

func Test_longestPrefix(t *testing.T) {
	type args struct {
		s string
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
				s: "ababab",
			},
			want: "",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := longestPrefix(tt.args.s); got != tt.want {
				t.Errorf("longestPrefix() = %v, want %v", got, tt.want)
			}
		})
	}
}
