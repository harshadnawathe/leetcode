package firstcompletelypaintedroworcolumn

import (
	"fmt"
	"reflect"
	"testing"
)

func TestFirstCompleteIndex(t *testing.T) {
	tests := []struct {
		arr  []int
		mat  [][]int
		want int
	}{
		{[]int{1, 3, 4, 2}, [][]int{{1, 4}, {2, 3}}, 2},
		{[]int{2, 8, 7, 4, 1, 3, 5, 6, 9}, [][]int{{3, 2, 5}, {1, 4, 6}, {8, 7, 9}}, 3},
	}

	for i, test := range tests {
		testName := fmt.Sprintf("Test_%v", i)
		t.Run(testName, func(t *testing.T) {
			got := firstCompleteIndex(
				test.arr,
				test.mat,
			)

			if !reflect.DeepEqual(got, test.want) {
				t.Errorf("Failed. want= %v got= %v", test.want, got)
			}
		})
	}
}
