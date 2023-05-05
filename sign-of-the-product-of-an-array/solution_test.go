package signoftheproductofanarray

import (
	"fmt"
	"reflect"
	"testing"
)

func TestArraySign(t *testing.T) {
	tests := []struct {
		nums []int
		want int
	}{
		{[]int{-1, -2, -3, -4, 3, 2, 1}, 1},
		{[]int{1, 5, 0, 2, -3}, 0},
		{[]int{-1, 1, -1, 1, -1}, -1},
	}

	for i, test := range tests {
		testName := fmt.Sprintf("Test_%v", i)
		t.Run(testName, func(t *testing.T) {
			got := arraySign(
				test.nums,
			)

			if !reflect.DeepEqual(got, test.want) {
				t.Errorf("Failed. want= %v got= %v", test.want, got)
			}
		})
	}
}
