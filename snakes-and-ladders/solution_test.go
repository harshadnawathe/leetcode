package snakeandladders

import "testing"

func Test_snakesAndLadders(t *testing.T) {
	type args struct {
		board [][]int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "Example 1",
			args: args{
				board: [][]int{
					{-1, -1, -1, -1, -1, -1},
					{-1, -1, -1, -1, -1, -1},
					{-1, -1, -1, -1, -1, -1},
					{-1, 35, -1, -1, 13, -1},
					{-1, -1, -1, -1, -1, -1},
					{-1, 15, -1, -1, -1, -1},
				},
			},
			want: 4,
		},
		{
			name: "Example 2",
			args: args{
				board: [][]int{
					{-1, -1},
					{1, 3},
				},
			},
			want: 1,
		},
		{
			name: "Example 3",
			args: args{
				board: [][]int{
					{1, 1, -1},
					{1, 1, 1},
					{-1, 1, 1},
				},
			},
			want: -1,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := snakesAndLadders(tt.args.board); got != tt.want {
				t.Errorf("snakesAndLadders() = %v, want %v", got, tt.want)
			}
		})
	}
}
