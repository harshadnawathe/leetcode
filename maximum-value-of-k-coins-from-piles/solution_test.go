package maximumvalueofkcoinsfrompiles

import (
	"fmt"
	"reflect"
	"testing"
)

func TestMaxValueOfCoins(t *testing.T) {
	tests := []struct {
		piles [][]int
		k     int
		want  int
	}{
		{[][]int{{1, 100, 3}, {7, 8, 9}}, 2, 101},
		{[][]int{{100}, {100}, {100}, {100}, {100}, {100}, {1, 1, 1, 1, 1, 1, 700}}, 7, 706},
	}

	for i, test := range tests {
		testName := fmt.Sprintf("Test_%v", i)
		t.Run(testName, func(t *testing.T) {
			got := maxValueOfCoins(
				test.piles,
				test.k,
			)

			if !reflect.DeepEqual(got, test.want) {
				t.Errorf("Failed. want= %v got= %v", test.want, got)
			}
		})
	}
}
