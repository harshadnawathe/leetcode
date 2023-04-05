package removeduplicatesfromsortedarray

import (
	"fmt"
	"reflect"
	"testing"
)

func TestRemoveDuplicates(t *testing.T) {

	tests := []struct {
		nums []int
		want int
	}{
		{[]int{1}, 1},
		{[]int{2, 2}, 1},
	}

	for i, test := range tests {
		testName := fmt.Sprintf("Test_%v", i)
		t.Run(testName, func(t *testing.T) {
			got := removeDuplicates(
				test.nums,
			)

			if !reflect.DeepEqual(got, test.want) {
				t.Errorf("Failed. want= %v got= %v", test.want, got)
				t.FailNow()
			}

			if containsDuplicates(test.nums[:got]) {
				t.Errorf("Failed. Duplicates not removed. got= %v", test.nums[:got])
			}
		})
	}
}

func containsDuplicates(nums []int) bool {
	m := make(map[int]struct{})
	for _, num := range nums {
		if _, found := m[num]; found {
			return true
		}
		m[num] = struct{}{}
	}

	return false
}
