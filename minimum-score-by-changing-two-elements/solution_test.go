package minimumscorebychangingtwoelements

import (
	"fmt"
	"reflect"
	"testing"
)

func TestMinimizeSum(t *testing.T) {

	tests := []struct {
		nums []int
		want int
	}{
		{[]int{1, 4, 3}, 0},
		{[]int{1, 4, 7, 8, 5}, 3},
	}

	for i, test := range tests {
		testName := fmt.Sprintf("Test_%v", i)
		t.Run(testName, func(t *testing.T) {
			got := minimizeSum(
				test.nums,
			)

			if !reflect.DeepEqual(got, test.want) {
				t.Errorf("Failed. want= %v got= %v", test.want, got)
			}
		})
	}
}
