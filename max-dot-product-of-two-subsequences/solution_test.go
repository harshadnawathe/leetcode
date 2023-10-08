package maxdotproductoftwosubsequences

import (
	"fmt"
	"reflect"
	"testing"
)

func TestMaxDotProduct(t *testing.T) {
	tests := []struct {
		nums1 []int
		nums2 []int
		want  int
	}{
		{
			nums1: []int{2, 1, -2, 5},
			nums2: []int{3, 0, -6},
			want:  18,
		},
		{
			nums1: []int{3, -2},
			nums2: []int{2, -6, 7},
			want:  21,
		},
		{
			nums1: []int{-1, -1},
			nums2: []int{1, 1},
			want:  -1,
		},
	}

	for i, test := range tests {
		testName := fmt.Sprintf("Test_%v", i)
		t.Run(testName, func(t *testing.T) {
			got := maxDotProduct(
				test.nums1,
				test.nums2,
			)

			if !reflect.DeepEqual(got, test.want) {
				t.Errorf("Failed. want= %v got= %v", test.want, got)
			}
		})
	}
}
