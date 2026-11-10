package findserversthathandledmostnumberofrequests

import (
	"reflect"
	"testing"
)

func Test_busiestServers(t *testing.T) {
	tests := []struct {
		name string // description of this test case
		// Named input parameters for target function.
		k       int
		arrival []int
		load    []int
		want    []int
	}{
		{
			name:    "Example 1",
			k:       3,
			arrival: []int{1, 2, 3, 4, 5},
			load:    []int{5, 2, 3, 3, 3},
			want:    []int{1},
		},
		{
			name:    "Example 2",
			k:       3,
			arrival: []int{1, 2, 3, 4},
			load:    []int{1, 2, 1, 2},
			want:    []int{0},
		},
		{
			name:    "Example 3",
			k:       3,
			arrival: []int{1, 2, 3},
			load:    []int{10, 12, 11},
			want:    []int{0, 1, 2},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := busiestServers(tt.k, tt.arrival, tt.load)
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("busiestServers() = %v, want %v", got, tt.want)
			}
		})
	}
}
