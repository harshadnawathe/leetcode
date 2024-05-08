package relativeranks

import (
	"reflect"
	"testing"
)

func Test_findRelativeRanks(t *testing.T) {
	type args struct {
		score []int
	}
	tests := []struct {
		name string
		args args
		want []string
	}{
		{
			name: "Example 1",
			args: args{
				score: []int{5, 4, 3, 2, 1},
			},
			want: []string{"Gold Medal", "Silver Medal", "Bronze Medal", "4", "5"},
		},
		{
			name: "Example 2",
			args: args{
				score: []int{10, 3, 8, 9, 4},
			},
			want: []string{"Gold Medal", "5", "Bronze Medal", "Silver Medal", "4"},
		},
		{
			name: "Example 3",
			args: args{
				score: []int{1, 2, 3, 4, 5},
			},
			want: []string{"5", "4", "Bronze Medal", "Silver Medal", "Gold Medal"},
		},
		{
			name: "Example 4",
			args: args{
				score: []int{1},
			},
			want: []string{"Gold Medal"},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := findRelativeRanks(tt.args.score); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("findRelativeRanks() = %v, want %v", got, tt.want)
			}
		})
	}
}
