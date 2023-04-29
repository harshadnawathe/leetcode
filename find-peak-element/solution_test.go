package findpeakelement

import (
	"fmt"
	"reflect"
	"testing"
)

func TestFindPeakElement(t *testing.T) {
	tests := []struct {
		nums []int
		want int
	}{
		{[]int{1, 2, 3, 1}, 2},
		{[]int{1, 2, 1, 3, 5, 6, 4}, 5},
		{[]int{1, 2, 3}, 2},
		{[]int{3, 2, 1}, 0},
	}

	for i, test := range tests {
		testName := fmt.Sprintf("Test_%v", i)
		t.Run(testName, func(t *testing.T) {
			got := findPeakElement(
				test.nums,
			)

			if !reflect.DeepEqual(got, test.want) {
				t.Errorf("Failed. want= %v got= %v", test.want, got)
			}
		})
	}
}
