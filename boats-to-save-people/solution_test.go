package boatstosavepeople

import (
	"fmt"
	"reflect"
	"testing"
)

func TestNumRescueBoats(t *testing.T) {

	tests := []struct {
		people []int
		limit  int
		want   int
	}{
		{[]int{1, 2}, 3, 1},
		{[]int{3, 2, 2, 1}, 3, 3},
		{[]int{3, 5, 3, 4}, 5, 4},
	}

	for i, test := range tests {
		testName := fmt.Sprintf("Test_%v", i)
		t.Run(testName, func(t *testing.T) {
			got := numRescueBoats(
				test.people,
				test.limit,
			)

			if !reflect.DeepEqual(got, test.want) {
				t.Errorf("Failed. want= %v got= %v", test.want, got)
			}
		})
	}
}
