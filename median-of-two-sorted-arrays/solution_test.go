package medianoftwosortedarrays

import (
	"fmt"
	"testing"
)

func TestFindMedianOfSrortedArrays(t *testing.T) {
	tests := []struct {
		nums1, nums2 []int
		want         float64
	}{
		{[]int{1, 3}, []int{2}, 2},
		{[]int{1, 2}, []int{3, 4}, 2.5},
	}

	for i, test := range tests {
		testName := fmt.Sprintf("Test_%v", i)
		t.Run(testName, func(t *testing.T) {
			got := findMedianSortedArrays(test.nums1, test.nums2)

			if got != test.want {
				t.Errorf("Failed. want= %v got= %v", test.want, got)
			}
		})
	}

}
