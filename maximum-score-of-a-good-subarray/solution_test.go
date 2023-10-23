package maximumscoreofagoodsubarray

import (
	"fmt"
	"reflect"
	"testing"
)

func TestMaximumScore(t *testing.T) {
	tests := []struct {
		nums []int
		k    int
		want int
	}{
		{[]int{1, 4, 3, 7, 4, 5}, 3, 15},
		{[]int{5, 5, 4, 5, 4, 1, 1, 1}, 0, 20},
	}

	for i, test := range tests {
		testName := fmt.Sprintf("Test_%v", i)
		t.Run(testName, func(t *testing.T) {
			got := maximumScore(
				test.nums,
				test.k,
			)

			if !reflect.DeepEqual(got, test.want) {
				t.Errorf("Failed. want= %v got= %v", test.want, got)
			}
		})
	}
}
