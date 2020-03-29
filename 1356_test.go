package leetcode

//
//func Test_getOneBitCount(t *testing.T) {
//	type args struct {
//		i int
//	}
//	tests := []struct {
//		name string
//		args args
//		want int
//	}{
//		// TODO: Add test cases.
//		{
//			name: "",
//			args: args{
//				i: 12312313,
//			},
//			want: 0,
//		},
//		{
//			name: "",
//			args: args{
//				i: 98765467890,
//			},
//			want: 0,
//		},
//		{
//			name: "",
//			args: args{
//				i: 34567897654,
//			},
//			want: 0,
//		},
//	}
//	for _, tt := range tests {
//		t.Run(tt.name, func(t *testing.T) {
//			if got := getOneBitCount(tt.args.i); got != strings.Count(strconv.FormatInt(int64(tt.args.i),2), "1") {
//				t.Errorf("getOneBitCount() = %v, want %v", got, tt.want)
//			}
//		})
//	}
//}
//
//func Test_sortByBits(t *testing.T) {
//	type args struct {
//		arr []int
//	}
//	tests := []struct {
//		name string
//		args args
//		want []int
//	}{
//		// TODO: Add test cases.
//		{
//			name: "",
//			args: args{
//				arr: []int{1111,7644,1107,6978,8742,1,7403,7694,9193,4401,377,8641,5311,624,3554,6631},
//			},
//			want: nil,
//		},
//	}
//	for _, tt := range tests {
//		t.Run(tt.name, func(t *testing.T) {
//			if got := sortByBits(tt.args.arr); !reflect.DeepEqual(got, tt.want) {
//				oneBitNum := make([]int,0)
//				for _, i:= range got {
//					oneBitNum = append(oneBitNum, getOneBitCount(i))
//				}
//				tools.Print(oneBitNum)
//				t.Errorf("sortByBits() = %v, want %v", got, tt.want)
//			}
//		})
//	}
//}
