package nqueens

import (
	"fmt"
	"testing"

	"github.com/google/go-cmp/cmp"
	"github.com/google/go-cmp/cmp/cmpopts"
)

var sortAscending = cmpopts.SortSlices(func(a, b []string) bool {
	return a[0] < b[0]
})

func TestSolveNQueens(t *testing.T) {
	tests := []struct {
		n    int
		want [][]string
	}{
		{1, [][]string{{"Q"}}},
		{4, [][]string{{".Q..", "...Q", "Q...", "..Q."}, {"..Q.", "Q...", "...Q", ".Q.."}}},
	}

	for i, test := range tests {
		testName := fmt.Sprintf("Test_%v", i)
		t.Run(testName, func(t *testing.T) {
			got := solveNQueens(
				test.n,
			)

			if !cmp.Equal(got, test.want, sortAscending) {
				t.Errorf("Failed. want= %v got= %v", test.want, got)
			}
		})
	}
}
