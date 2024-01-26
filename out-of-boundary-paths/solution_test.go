package outofboundarypaths

import "testing"

func Test_findPaths(t *testing.T) {
	type args struct {
		m           int
		n           int
		maxMove     int
		startRow    int
		startColumn int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{"Zero moves", args{1, 1, 0, 0, 0}, 0},
		{"Single cell, single move", args{1, 1, 1, 0, 0}, 4},
		{"Example 1", args{2, 2, 2, 0, 0}, 6},
		{"Example 2", args{1, 3, 3, 0, 1}, 12},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := findPaths(tt.args.m, tt.args.n, tt.args.maxMove, tt.args.startRow, tt.args.startColumn); got != tt.want {
				t.Errorf("findPaths() = %v, want %v", got, tt.want)
			}
		})
	}
}
