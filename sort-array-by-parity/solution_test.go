package sortarraybyparity

import (
	"fmt"
	"testing"
)

func TestSortArrayByParity(t *testing.T) {
	tests := []struct {
		nums []int
	}{
		{[]int{0}},
		{[]int{1, 3}},
		{[]int{0, 2}},
		{[]int{3, 1, 2, 4}},
		{[]int{1, 2, 3, 4}},
	}

	for i, test := range tests {
		testName := fmt.Sprintf("Test_%v", i)
		t.Run(testName, func(t *testing.T) {
			got := sortArrayByParity(
				test.nums,
			)
			t.Logf("Got= %v", got)
			if !isPartitioned(got) {
				t.Errorf("Failed. want isPartitioned got= %v", got)
			}
		})
	}
}

func isPartitioned(nums []int) bool {
	pp := len(nums)
	for i, num := range nums {
		if num&0x1 == 1 {
			pp = i
			break
		}
	}

	for _, num := range nums[pp:] {
		if num&0x1 == 0 {
			return false
		}
	}

	return true
}
