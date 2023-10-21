package constrainedsubsequencesum

import (
	"fmt"
	"reflect"
	"testing"
)

func TestConstrainedSubsetSum(t *testing.T) {
	tests := []struct {
		nums []int
		k    int
		want int
	}{
		{[]int{10, 2, -10, 5, 20}, 2, 37},
		{[]int{-1, -2, -3}, 1, -1},
		{[]int{10, -2, -10, -5, 20}, 2, 23},
	}

	for i, test := range tests {
		testName := fmt.Sprintf("Test_%v", i)
		t.Run(testName, func(t *testing.T) {
			got := constrainedSubsetSum(
				test.nums,
				test.k,
			)

			if !reflect.DeepEqual(got, test.want) {
				t.Errorf("Failed. want= %v got= %v", test.want, got)
			}
		})
	}
}
