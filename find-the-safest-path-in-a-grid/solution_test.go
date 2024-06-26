package findthesafestpathinagrid

import "testing"

func Test_maximumSafenessFactor(t *testing.T) {
	type args struct {
		grid [][]int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "Example 1",
			args: args{
				grid: [][]int{
					{1, 0, 0},
					{0, 0, 0},
					{0, 0, 1},
				},
			},
			want: 0,
		},
		{
			name: "Example 2",
			args: args{
				grid: [][]int{
					{0, 0, 1}, {0, 0, 0}, {0, 0, 0},
				},
			},
			want: 2,
		},
		{
			name: "Example 3",
			args: args{
				grid: [][]int{
					{0, 0, 0, 1},
					{0, 0, 0, 0},
					{0, 0, 0, 0},
					{1, 0, 0, 0},
				},
			},
			want: 2,
		},
		{
			name: "Example 4",
			args: args{
				grid: [][]int{
					{0, 1, 1},
					{0, 1, 1},
					{1, 1, 0},
				},
			},
			want: 0,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := maximumSafenessFactor(tt.args.grid); got != tt.want {
				t.Errorf("maximumSafenessFactor() = %v, want %v", got, tt.want)
			}
		})
	}
}
