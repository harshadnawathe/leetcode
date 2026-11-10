package climbingstairsii

import (
	"testing"
)

func Test_climbStairs(t *testing.T) {
	tests := []struct {
		name  string
		n     int
		costs []int
		want  int
	}{
		{
			name:  "Example 1",
			n:     4,
			costs: []int{1, 2, 3, 4},
			want:  13,
		},
		{
			name:  "Example 2",
			n:     4,
			costs: []int{5, 1, 6, 2},
			want:  11,
		},
		{
			name:  "Example 3",
			n:     3,
			costs: []int{9, 8, 3},
			want:  12,
		},
		{
			name:  "Example 4",
			n:     1,
			costs: []int{9},
			want:  10,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := climbStairs(tt.n, tt.costs)
			// TODO: update the condition below to compare got with tt.want.
			if got != tt.want {
				t.Errorf("climbStairs() = %v, want %v", got, tt.want)
			}
		})
	}
}
