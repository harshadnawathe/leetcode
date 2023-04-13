package firstmissingpositive

import (
	"fmt"
	"reflect"
	"testing"
)

func TestFirstMissingPositive(t *testing.T) {

	tests := []struct {
		nums []int
		want int
	}{
		{[]int{1, 2, 0}, 3},
		{[]int{3, 4, -1, 1}, 2},
		{[]int{7, 8, 9, 11, 12}, 1},
		{[]int{1}, 2},
		{[]int{3, 3, 1, 4, 0}, 2},
		{[]int{0, 1, 2}, 3},
	}

	for i, test := range tests {
		testName := fmt.Sprintf("Test_%v", i)
		t.Run(testName, func(t *testing.T) {
			got := firstMissingPositive(
				test.nums,
			)

			if !reflect.DeepEqual(got, test.want) {
				t.Errorf("Failed. want= %v got= %v", test.want, got)
			}
		})
	}
}
