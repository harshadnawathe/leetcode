package uncrossedlines

import (
	"fmt"
	"reflect"
	"testing"
)

func TestMaxUncrossedLines(t *testing.T) {
	tests := []struct {
		nums1 []int
		nums2 []int
		want  int
	}{
		{[]int{1, 4, 2}, []int{1, 2, 4}, 2},
		{[]int{2, 5, 1, 2, 5}, []int{10, 5, 2, 1, 5, 2}, 3},
		{[]int{1, 3, 7, 1, 7, 5}, []int{1, 9, 2, 5, 1}, 2},
	}

	for i, test := range tests {
		testName := fmt.Sprintf("Test_%v", i)
		t.Run(testName, func(t *testing.T) {
			got := maxUncrossedLines(
				test.nums1,
				test.nums2,
			)

			if !reflect.DeepEqual(got, test.want) {
				t.Errorf("Failed. want= %v got= %v", test.want, got)
			}
		})
	}
}
