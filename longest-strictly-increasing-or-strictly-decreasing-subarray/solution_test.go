package longeststrictlyincreasingorstrictlydecreasingsubarray

import (
	"reflect"
	"testing"
)

func Test_longestMonotonicSubarray(t *testing.T) {
	tests := []struct {
		name string
		nums []int
		want int
	}{
		{
			name: "Example 1",
			nums: []int{1, 4, 3, 3, 2},
			want: 2,
		},
		{
			name: "Example 2",
			nums: []int{3, 3, 3, 3},
			want: 1,
		},
		{
			name: "Example 3",
			nums: []int{3, 2, 1},
			want: 3,
		},
		{
			name: "Example 4",
			nums: []int{1},
			want: 1,
		},
		{
			name: "Example 5",
			nums: []int{1,10,10},
			want: 2,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := longestMonotonicSubarray(tt.nums)
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("longestMonotonicSubarray() = %v, want %v", got, tt.want)
			}
		})
	}
}
