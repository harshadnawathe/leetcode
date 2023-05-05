package createmaximumnumber

import (
	"fmt"
	"reflect"
	"testing"
)

func TestMaxNumber(t *testing.T) {
	tests := []struct {
		nums1 []int
		nums2 []int
		k     int
		want  []int
	}{
		{[]int{3, 4, 6, 5}, []int{9, 1, 2, 5, 8, 3}, 5, []int{9, 8, 6, 5, 3}},
		{[]int{6, 7}, []int{6, 0, 4}, 5, []int{6, 7, 6, 0, 4}},
		{[]int{3, 9}, []int{8, 9}, 3, []int{9, 8, 9}},
		{[]int{2, 5, 6, 4, 4, 0}, []int{7, 3, 8, 0, 6, 5, 7, 6, 2}, 15, []int{7, 3, 8, 2, 5, 6, 4, 4, 0, 6, 5, 7, 6, 2, 0}},
		{[]int{8, 1, 8, 8, 6}, []int{4}, 2, []int{8, 8}},
	}

	for i, test := range tests {
		testName := fmt.Sprintf("Test_%v", i)
		t.Run(testName, func(t *testing.T) {
			got := maxNumber(
				test.nums1,
				test.nums2,
				test.k,
			)

			if !reflect.DeepEqual(got, test.want) {
				t.Errorf("Failed. want= %v got= %v", test.want, got)
			}
		})
	}
}
