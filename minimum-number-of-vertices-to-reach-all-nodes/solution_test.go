package minimumnumberofverticestoreachallnodes

import (
	"fmt"
	"testing"

	"github.com/google/go-cmp/cmp"
	"github.com/google/go-cmp/cmp/cmpopts"
)

var sortAscending = cmpopts.SortSlices(func(a, b int) bool { return a < b })

func TestFindSmallestSetOfVertices(t *testing.T) {
	tests := []struct {
		n     int
		edges [][]int
		want  []int
	}{
		{6, [][]int{{0, 1}, {0, 2}, {2, 5}, {3, 4}, {4, 2}}, []int{0, 3}},
		{5, [][]int{{0, 1}, {2, 1}, {3, 1}, {1, 4}, {2, 4}}, []int{0, 2, 3}},
	}

	for i, test := range tests {
		testName := fmt.Sprintf("Test_%v", i)
		t.Run(testName, func(t *testing.T) {
			got := findSmallestSetOfVertices(
				test.n,
				test.edges,
			)

			if !cmp.Equal(got, test.want, sortAscending) {
				t.Errorf("Failed. want= %v got= %v", test.want, got)
			}
		})
	}
}
