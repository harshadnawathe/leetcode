package maximizescoreafternoperations

import (
	"fmt"
	"reflect"
	"testing"
)

func TestMaxScore(t *testing.T) {
	tests := []struct {
		nums []int
		want int
	}{
		{[]int{1, 2}, 1},
		{[]int{3, 4, 6, 8}, 11},
		{[]int{1, 2, 3, 4, 5, 6}, 14},
	}

	for i, test := range tests {
		testName := fmt.Sprintf("Test_%v", i)
		t.Run(testName, func(t *testing.T) {
			got := maxScore(
				test.nums,
			)

			if !reflect.DeepEqual(got, test.want) {
				t.Errorf("Failed. want= %v got= %v", test.want, got)
			}
		})
	}
}
