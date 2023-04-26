package describethepainting

import (
	"fmt"
	"reflect"
	"testing"
)

func TestSplitPainting(t *testing.T) {
	tests := []struct {
		segments [][]int
		want     [][]int64
	}{
		{[][]int{{1, 4, 5}}, [][]int64{{1, 4, 5}}},
		{[][]int{{1, 4, 5}, {4, 7, 7}}, [][]int64{{1, 4, 5}, {4, 7, 7}}},
		{[][]int{{1, 4, 5}, {4, 7, 7}, {1, 7, 9}}, [][]int64{{1, 4, 14}, {4, 7, 16}}},
		{[][]int{{1, 7, 9}, {6, 8, 15}, {8, 10, 7}}, [][]int64{{1, 6, 9}, {6, 7, 24}, {7, 8, 15}, {8, 10, 7}}},
		{[][]int{{1, 4, 5}, {1, 4, 7}, {4, 7, 1}, {4, 7, 11}}, [][]int64{{1, 4, 12}, {4, 7, 12}}},
	}

	for i, test := range tests {
		testName := fmt.Sprintf("Test_%v", i)
		t.Run(testName, func(t *testing.T) {
			got := splitPainting(
				test.segments,
			)

			if !reflect.DeepEqual(got, test.want) {
				t.Errorf("Failed. want= %v got= %v", test.want, got)
			}
		})
	}
}
