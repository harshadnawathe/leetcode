package findplayerswithzerooronelosses

import (
	"reflect"
	"testing"
)

func Test_findWinners(t *testing.T) {
	type args struct {
		matches [][]int
	}
	tests := []struct {
		name string
		args args
		want [][]int
	}{
		{
			"Example 1",
			args{matches: [][]int{{1, 3}, {2, 3}, {3, 6}, {5, 6}, {5, 7}, {4, 5}, {4, 8}, {4, 9}, {10, 4}, {10, 9}}},
			[][]int{{1, 2, 10}, {4, 5, 7, 8}},
		},
		{
			"Example 2",
			args{matches: [][]int{{2, 3}, {1, 3}, {5, 4}, {6, 4}}},
			[][]int{{1, 2, 5, 6}, {}},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := findWinners(tt.args.matches)

			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("findWinners() = %v, want %v", got, tt.want)
			}
		})
	}
}
