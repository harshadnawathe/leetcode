package matrixdiagonalsum

import (
	"fmt"
	"reflect"
	"testing"
)

func TestDiagonalSum(t *testing.T) {
	tests := []struct {
		mat  [][]int
		want int
	}{
		{[][]int{{1, 2, 3}, {4, 5, 6}, {7, 8, 9}}, 25},
		{[][]int{{1, 1, 1, 1}, {1, 1, 1, 1}, {1, 1, 1, 1}, {1, 1, 1, 1}}, 8},
		{[][]int{{5}}, 5},
	}

	for i, test := range tests {
		testName := fmt.Sprintf("Test_%v", i)
		t.Run(testName, func(t *testing.T) {
			got := diagonalSum(
				test.mat,
			)

			if !reflect.DeepEqual(got, test.want) {
				t.Errorf("Failed. want= %v got= %v", test.want, got)
			}
		})
	}
}
