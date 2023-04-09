package nextpermutation

import (
	"fmt"
	"reflect"
	"testing"
)

func TestSolution(t *testing.T) {

	tests := []struct {
		nums []int
		want []int
	}{
		{[]int{1, 2, 3, 4}, []int{1, 2, 4, 3}},
		{[]int{1, 2, 4, 3}, []int{1, 3, 2, 4}},
		{[]int{1, 3, 2, 4}, []int{1, 3, 4, 2}},
		{[]int{1, 3, 4, 2}, []int{1, 4, 2, 3}},
		{[]int{1, 4, 2, 3}, []int{1, 4, 3, 2}},
		{[]int{1, 4, 3, 2}, []int{2, 1, 3, 4}},
		{[]int{4, 3, 2, 1}, []int{1, 2, 3, 4}},

		{[]int{1, 2, 3}, []int{1, 3, 2}},
		{[]int{3, 2, 1}, []int{1, 2, 3}},
		{[]int{1, 1, 5}, []int{1, 5, 1}},
	}

	for i, test := range tests {
		testName := fmt.Sprintf("Test_%v", i)
		t.Run(testName, func(t *testing.T) {
			nextPermutation(test.nums)

			if !reflect.DeepEqual(test.nums, test.want) {
				t.Errorf("Failed. want= %v got= %v", test.want, test.nums)
			}
		})
	}
}
