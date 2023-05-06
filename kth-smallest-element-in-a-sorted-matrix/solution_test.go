package kthsmallestelementinasortedmatrix

import (
	"fmt"
	"reflect"
	"testing"
)

func TestKthSmallest(t *testing.T) {
	tests := []struct {
		matrix [][]int
		k      int
		want   int
	}{
		{[][]int{{1, 5, 9}, {10, 11, 13}, {12, 13, 15}}, 8, 13},
		{[][]int{{-5}}, 1, -5},
	}

	for i, test := range tests {
		testName := fmt.Sprintf("Test_%v", i)
		t.Run(testName, func(t *testing.T) {
			got := kthSmallest(
				test.matrix,
				test.k,
			)

			if !reflect.DeepEqual(got, test.want) {
				t.Errorf("Failed. want= %v got= %v", test.want, got)
			}
		})
	}
}
