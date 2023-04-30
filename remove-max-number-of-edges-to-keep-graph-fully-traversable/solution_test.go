package removemaxnumberofedgestokeepgraphfullytraversable

import (
	"fmt"
	"reflect"
	"testing"
)

func TestMaxNumEdgesToRemove(t *testing.T) {
	tests := []struct {
		n     int
		edges [][]int
		want  int
	}{
		{4, [][]int{{3, 1, 2}, {3, 2, 3}, {1, 1, 3}, {1, 2, 4}, {1, 1, 2}, {2, 3, 4}}, 2},
		{4, [][]int{{3, 1, 2}, {3, 2, 3}, {1, 1, 4}, {2, 1, 4}}, 0},
		{4, [][]int{{3, 2, 3}, {1, 1, 2}, {2, 3, 4}}, -1},
	}

	for i, test := range tests {
		testName := fmt.Sprintf("Test_%v", i)
		t.Run(testName, func(t *testing.T) {
			got := maxNumEdgesToRemove(
				test.n,
				test.edges,
			)

			if !reflect.DeepEqual(got, test.want) {
				t.Errorf("Failed. want= %v got= %v", test.want, got)
			}
		})
	}
}
