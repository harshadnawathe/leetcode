package reducingdishes

import (
	"fmt"
	"reflect"
	"testing"
)

func TestMaxSatisfaction(t *testing.T) {

	tests := []struct {
		satisfaction []int
		want         int
	}{
		{[]int{-1, -8, 0, 5, -9}, 14},
		{[]int{4, 3, 2}, 20},
		{[]int{-1, -4, -5}, 0},
	}

	for i, test := range tests {
		testName := fmt.Sprintf("Test_%v", i)
		t.Run(testName, func(t *testing.T) {
			got := maxSatisfaction(
				test.satisfaction,
			)

			if !reflect.DeepEqual(got, test.want) {
				t.Errorf("Failed. want= %v got= %v", test.want, got)
			}
		})
	}
}
