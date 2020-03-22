package leetcode

import "testing"

func Test_numOfMinutes(t *testing.T) {
	type args struct {
		n          int
		headID     int
		manager    []int
		informTime []int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "",
			args: args{
				n:          11,
				headID:     4,
				manager:    []int{5, 9, 6, 10, -1, 8, 9, 1, 9, 3, 4},
				informTime: []int{0, 213, 0, 253, 686, 170, 975, 0, 261, 309, 337},
			},
			want: 0,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := numOfMinutes(tt.args.n, tt.args.headID, tt.args.manager, tt.args.informTime); got != tt.want {
				t.Errorf("numOfMinutes() = %v, want %v", got, tt.want)
			}
		})
	}
}
