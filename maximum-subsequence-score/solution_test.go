package maximumsubsequencescore

import (
	"fmt"
	"reflect"
	"testing"
)

func TestMaxScore(t *testing.T) {
	tests := []struct {
		nums1 []int
		nums2 []int
		k     int
		want  int64
	}{
		{[]int{1, 3, 3, 2}, []int{2, 1, 3, 4}, 3, 12},
		{[]int{4, 2, 3, 1, 1}, []int{7, 5, 10, 9, 6}, 1, 30},
	}

	for i, test := range tests {
		testName := fmt.Sprintf("Test_%v", i)
		t.Run(testName, func(t *testing.T) {
			got := maxScore(
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
