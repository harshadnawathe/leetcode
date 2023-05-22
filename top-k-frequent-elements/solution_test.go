package topkfrequentelements

import (
	"fmt"
	"testing"

	"github.com/google/go-cmp/cmp"
	"github.com/google/go-cmp/cmp/cmpopts"
)

var sortAscending = cmpopts.SortSlices(func(i, j int) bool {
	return i < j
})

func TestTopKFrequent(t *testing.T) {
	tests := []struct {
		nums []int
		k    int
		want []int
	}{
		{[]int{1, 1, 1, 2, 2, 3}, 2, []int{1, 2}},
		{[]int{1}, 1, []int{1}},
		{[]int{3, 0, 1, 0}, 1, []int{0}},
	}

	for i, test := range tests {
		testName := fmt.Sprintf("Test_%v", i)
		t.Run(testName, func(t *testing.T) {
			got := topKFrequent(
				test.nums,
				test.k,
			)

			if !cmp.Equal(got, test.want, sortAscending) {
				t.Errorf("Failed. want= %v got= %v", test.want, got)
			}
		})
	}
}
