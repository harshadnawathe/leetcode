package largestcolorvalueinadirectedgraph

import (
	"fmt"
	"reflect"
	"testing"
)

func TestLargestPathValue(t *testing.T) {

	tests := []struct {
		colors string
		edges  [][]int
		want   int
	}{
		{"abaca", [][]int{{0, 1}, {0, 2}, {2, 3}, {3, 4}}, 3},
		{"a", [][]int{{0, 0}}, -1},
		{"g", [][]int{}, 1},
	}

	for i, test := range tests {
		testName := fmt.Sprintf("Test_%v", i)
		t.Run(testName, func(t *testing.T) {
			got := largestPathValue(
				test.colors,
				test.edges,
			)

			if !reflect.DeepEqual(got, test.want) {
				t.Errorf("Failed. want= %v got= %v", test.want, got)
			}
		})
	}
}
