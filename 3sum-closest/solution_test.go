package sumclosest

import (
	"fmt"
	"reflect"
	"testing"
)

func TestThreeSumClosest(t *testing.T) {

	tests := []struct {
		nums   []int
		target int
		want   int
	}{
		{[]int{-1, 2, 1, -4}, 1, 2},
		{[]int{0, 0, 0}, 1, 0},
	}

	for i, test := range tests {
		testName := fmt.Sprintf("Test_%v", i)
		t.Run(testName, func(t *testing.T) {
			got := threeSumClosest(
				test.nums,
				test.target,
			)

			if !reflect.DeepEqual(got, test.want) {
				t.Errorf("Failed. want= %v got= %v", test.want, got)
			}
		})
	}
}
