package reversepairs

import "testing"

func Test_reversePairs(t *testing.T) {
	type args struct {
		nums []int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "Example 1",
			args: args{
				nums: []int{1, 3, 2, 3, 1},
			},
			want: 2,
		},
		{
			name: "Example 2",
			args: args{
				nums: []int{2, 4, 3, 5, 1},
			},
			want: 3,
		},
		{
			name: "Example 3",
			args: args{
				nums: []int{2147483647, 2147483647, 2147483647, 2147483647, 2147483647, 2147483647},
			},
			want: 0,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := reversePairs(tt.args.nums); got != tt.want {
				t.Errorf("reversePairs() = %v, want %v", got, tt.want)
			}
		})
	}
}
