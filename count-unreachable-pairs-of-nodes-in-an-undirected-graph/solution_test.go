package countunreachablepairsofnodesinanundirectedgraph

import (
	"fmt"
	"reflect"
	"testing"
)

func TestCountPairs(t *testing.T) {
	tests := []struct {
		n     int
		edges [][]int
		want  int64
	}{
		{3, [][]int{{0, 1}, {0, 2}, {1, 2}}, 0},
		{7, [][]int{{0, 2}, {0, 5}, {2, 4}, {1, 6}, {5, 4}}, 14},
	}

	for i, test := range tests {
		testName := fmt.Sprintf("Test_%v", i)
		t.Run(testName, func(t *testing.T) {

			got := countPairs(test.n, test.edges)

			if !reflect.DeepEqual(got, test.want) {
				t.Errorf("Failed. want= %v got= %v", test.want, got)
			}
		})
	}
}
