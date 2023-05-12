package solvingquestionswithbrainpower

import (
	"fmt"
	"reflect"
	"testing"
)

func TestMostPoints(t *testing.T) {
	tests := []struct {
		questions [][]int
		want      int64
	}{
		{[][]int{{3, 2}, {4, 3}, {4, 4}, {2, 5}}, 5},
		{[][]int{{1, 1}, {2, 2}, {3, 3}, {4, 4}, {5, 5}}, 7},
	}

	for i, test := range tests {
		testName := fmt.Sprintf("Test_%v", i)
		t.Run(testName, func(t *testing.T) {
			got := mostPoints(
				test.questions,
			)

			if !reflect.DeepEqual(got, test.want) {
				t.Errorf("Failed. want= %v got= %v", test.want, got)
			}
		})
	}
}
