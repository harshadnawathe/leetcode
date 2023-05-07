package findthedifferenceoftwoarrays

import (
	"fmt"
	"reflect"
	"sort"
	"testing"
)

func TestFindDifference(t *testing.T) {
	tests := []struct {
		nums1 []int
		nums2 []int
		want  [][]int
	}{
		{[]int{1, 2, 3}, []int{2, 4, 6}, [][]int{{1, 3}, {4, 6}}},
		{[]int{1, 2, 3, 3}, []int{1, 1, 2, 2}, [][]int{{3}, {}}},
	}

	for i, test := range tests {
		testName := fmt.Sprintf("Test_%v", i)
		t.Run(testName, func(t *testing.T) {
			got := findDifference(
				test.nums1,
				test.nums2,
			)

			sort.Ints(got[0])
			sort.Ints(got[1])
			sort.Ints(test.want[0])
			sort.Ints(test.want[1])

			if !reflect.DeepEqual(got, test.want) {
				t.Errorf("Failed. want= %v got= %v", test.want, got)
			}
		})
	}
}
