package twosum

import (
	"fmt"
	"testing"

	"github.com/google/go-cmp/cmp"
	"github.com/google/go-cmp/cmp/cmpopts"
)

var sortAscending = cmpopts.SortSlices(func(a, b int) bool {
	return a < b
})

func TestTwoSum(t *testing.T) {
	tests := []struct {
		nums   []int
		target int
		want   []int
	}{
		{[]int{2, 7, 11, 15}, 9, []int{0, 1}},
		{[]int{3, 2, 4}, 6, []int{1, 2}},
		{[]int{3, 3}, 6, []int{0, 1}},
	}

	for i, test := range tests {
		testName := fmt.Sprintf("Test_%d", i)
		t.Run(testName, func(t *testing.T) {
			got := twoSum(test.nums, test.target)

			if !cmp.Equal(got, test.want, sortAscending) {
				t.Errorf("Failed. want= %v got= %v", test.want, got)
			}
		})
	}
}
