package removeelement

import (
	"fmt"
	"reflect"
	"testing"
)

func TestRemoveElement(t *testing.T) {

	tests := []struct {
		nums []int
		val  int
		want int
	}{
		{[]int{3, 2, 2, 3}, 3, 2},
		{[]int{0, 1, 2, 2, 3, 0, 4, 2}, 2, 5},
	}

	for i, test := range tests {
		testName := fmt.Sprintf("Test_%v", i)
		t.Run(testName, func(t *testing.T) {
			got := removeElement(
				test.nums,
				test.val,
			)

			if !reflect.DeepEqual(got, test.want) {
				t.Errorf("Failed. want= %v got= %v", test.want, got)
				t.FailNow()
			}

			if contains(test.nums[:got], test.val) {
				t.Errorf("Failed. val should be removed.")
			}
		})
	}
}

func contains(nums []int, k int) bool {
	for _, num := range nums {
		if num == k {
			return true
		}
	}
	return false
}
