package dividearrayinsetsofkconsecutivenumbers

import "testing"

func Test_isPossibleDivide(t *testing.T) {
	type args struct {
		nums []int
		k    int
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			name: "Example 1",
			args: args{
				nums: []int{1, 2, 3, 3, 4, 4, 5, 6},
				k:    4,
			},
			want: true,
		},
		{
			name: "Example 2",
			args: args{
				nums: []int{3, 2, 1, 2, 3, 4, 3, 4, 5, 9, 10, 11},
				k:    3,
			},
			want: true,
		},
		{
			name: "Example 3",
			args: args{
				nums: []int{1, 2, 3, 4},
				k:    3,
			},
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := isPossibleDivide(tt.args.nums, tt.args.k); got != tt.want {
				t.Errorf("isPossibleDivide() = %v, want %v", got, tt.want)
			}
		})
	}
}
