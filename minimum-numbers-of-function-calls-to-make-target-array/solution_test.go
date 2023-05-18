package minimumnumbersoffunctioncallstomaketargetarray

import (
	"fmt"
	"reflect"
	"testing"
)

func TestMinOperations(t *testing.T) {
	tests := []struct {
		nums []int
		want int
	}{
		{[]int{1, 5}, 5},
		{[]int{2, 2}, 3},
		{[]int{4, 2, 5}, 6},
		{[]int{0}, 0},
		{[]int{0, 0}, 0},
	}

	for i, test := range tests {
		testName := fmt.Sprintf("Test_%v", i)
		t.Run(testName, func(t *testing.T) {
			got := minOperations(
				test.nums,
			)

			if !reflect.DeepEqual(got, test.want) {
				t.Errorf("Failed. want= %v got= %v", test.want, got)
			}
		})
	}
}
