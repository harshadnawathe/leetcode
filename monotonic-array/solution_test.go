package monotonicarray

import (
	"fmt"
	"reflect"
	"testing"
)

func TestIsMonotonic(t *testing.T) {
	tests := []struct {
		nums []int
		want bool
	}{
		{[]int{1, 2, 2, 3}, true},
		{[]int{6, 5, 4, 4}, true},
		{[]int{1, 3, 2}, false},
		{[]int{1, 3, 2}, false},
		{[]int{1, 1, 2, 3}, true},
		{[]int{3, 3, 3, 1}, true},
		{[]int{1}, true},
	}

	for i, test := range tests {
		testName := fmt.Sprintf("Test_%v", i)
		t.Run(testName, func(t *testing.T) {
			got := isMonotonic(
				test.nums,
			)

			if !reflect.DeepEqual(got, test.want) {
				t.Errorf("Failed. want= %v got= %v", test.want, got)
			}
		})
	}
}
