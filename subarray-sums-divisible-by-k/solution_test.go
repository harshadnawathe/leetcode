package subarraysumsdivisiblebyk

import "testing"

func Test_subarraysDivByK(t *testing.T) {
	type args struct {
		nums []int
		k    int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "Example 1",
			args: args{
				nums: []int{4, 5, 0, -2, -3, 1},
				k:    5,
			},
			want: 7,
		},
		{
			name: "Example 2",
			args: args{
				nums: []int{5},
				k:    9,
			},
			want: 0,
		},
		{
			name: "Example 3",
			args: args{
				nums: []int{9},
				k:    9,
			},
			want: 1,
		},
		{
			name: "Example 4",
			args: args{
				nums: []int{-1, 2, 9},
				k:    2,
			},
			want: 2,
		},
		{
			name: "Example 5",
			args: args{
				nums: []int{2, -2, 2, -4},
				k:    6,
			},
			want: 2,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := subarraysDivByK(tt.args.nums, tt.args.k); got != tt.want {
				t.Errorf("subarraysDivByK() = %v, want %v", got, tt.want)
			}
		})
	}
}
