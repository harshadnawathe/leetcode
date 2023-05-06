package numberofsubsequencesthatsatisfythegivensumcondition

import (
	"fmt"
	"reflect"
	"testing"
)

func TestNumSubseq(t *testing.T) {
	tests := []struct {
		nums   []int
		target int
		want   int
	}{
		{[]int{3, 5, 6, 7}, 9, 4},
		{[]int{3, 3, 6, 8}, 10, 6},
		{[]int{2, 3, 3, 4, 6, 7}, 12, 61},
	}

	for i, test := range tests {
		testName := fmt.Sprintf("Test_%v", i)
		t.Run(testName, func(t *testing.T) {
			got := numSubseq(
				test.nums,
				test.target,
			)

			if !reflect.DeepEqual(got, test.want) {
				t.Errorf("Failed. want= %v got= %v", test.want, got)
			}
		})
	}
}
