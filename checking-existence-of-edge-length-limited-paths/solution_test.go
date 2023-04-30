package checkingexistenceofedgelengthlimitedpaths

import (
	"fmt"
	"reflect"
	"testing"
)

func TestDistanceLimitedPathsExist(t *testing.T) {
	tests := []struct {
		n        int
		edgeList [][]int
		queries  [][]int
		want     []bool
	}{
		{3, [][]int{{0, 1, 2}, {1, 2, 4}, {2, 0, 8}, {1, 0, 16}}, [][]int{{0, 1, 2}, {0, 2, 5}}, []bool{false, true}},
		{5, [][]int{{0, 1, 10}, {1, 2, 5}, {2, 3, 9}, {3, 4, 13}}, [][]int{{0, 4, 14}, {1, 4, 13}}, []bool{true, false}},
	}

	for i, test := range tests {
		testName := fmt.Sprintf("Test_%v", i)
		t.Run(testName, func(t *testing.T) {
			got := distanceLimitedPathsExist(
				test.n,
				test.edgeList,
				test.queries,
			)

			if !reflect.DeepEqual(got, test.want) {
				t.Errorf("Failed. want= %v got= %v", test.want, got)
			}
		})
	}
}
