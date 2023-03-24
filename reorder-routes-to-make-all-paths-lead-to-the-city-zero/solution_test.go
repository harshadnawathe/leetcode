package reorderroutestomakeallpathsleadtothecityzero

import (
	"fmt"
	"reflect"
	"testing"
)

func TestMinReorder(t *testing.T) {
	tests := []struct {
		n           int
		connections [][]int
		want        int
	}{
		{6, [][]int{{0, 1}, {1, 3}, {2, 3}, {4, 0}, {4, 5}}, 3},
		{5, [][]int{{1, 0}, {1, 2}, {3, 2}, {3, 4}}, 2},
		{3, [][]int{{1, 0}, {2, 0}}, 0},
		{4, [][]int{{0, 1}, {2, 0}, {3, 2}}, 1},
	}

	for i, test := range tests {
		testName := fmt.Sprintf("Test_%v", i)
		t.Run(testName, func(t *testing.T) {

			got := minReorder(test.n, test.connections)

			if !reflect.DeepEqual(got, test.want) {
				t.Errorf("Failed. want= %v got= %v", test.want, got)
			}
		})
	}
}
