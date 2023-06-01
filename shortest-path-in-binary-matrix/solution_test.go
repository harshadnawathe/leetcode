package shortestpathinbinarymatrix

import (
	"fmt"
	"reflect"
	"testing"
)

func TestShortestPathBinaryMatrix(t *testing.T) {
	tests := []struct {
		grid [][]int
		want int
	}{
		{[][]int{{0, 1}, {1, 0}}, 2},
		{[][]int{{0, 0, 0}, {1, 1, 0}, {1, 1, 0}}, 4},
		{[][]int{{1, 0, 0}, {1, 1, 0}, {1, 1, 0}}, -1},
		{[][]int{
			{0, 1, 1, 0, 0, 0},
			{0, 1, 0, 1, 1, 0},
			{0, 1, 1, 0, 1, 0},
			{0, 0, 0, 1, 1, 0},
			{1, 1, 1, 1, 1, 0},
			{1, 1, 1, 1, 1, 0},
		}, 14},
	}

	for i, test := range tests {
		testName := fmt.Sprintf("Test_%v", i)
		t.Run(testName, func(t *testing.T) {
			got := shortestPathBinaryMatrix(
				test.grid,
			)

			if !reflect.DeepEqual(got, test.want) {
				t.Errorf("Failed. want= %v got= %v", test.want, got)
			}
		})
	}
}
