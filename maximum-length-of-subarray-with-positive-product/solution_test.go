package maximumlengthofsubarraywithpositiveproduct

import (
	"fmt"
	"reflect"
	"testing"
)

func TestGetMaxLen(t *testing.T) {

	tests := []struct {
		nums []int
		want int
	}{
		//{[]int{}, 0}, // not valid input
		{[]int{1}, 1},
		{[]int{0}, 0},
		{[]int{1, -2}, 1},
		{[]int{1, -2, -3, 4}, 4},
		{[]int{0, 1, -2, -3, -4}, 3},
		{[]int{-1, -2, -3, 0, 1}, 2},
	}

	for i, test := range tests {
		testName := fmt.Sprintf("Test_%v", i)
		t.Run(testName, func(t *testing.T) {
			got := getMaxLen(
				test.nums,
			)

			if !reflect.DeepEqual(got, test.want) {
				t.Errorf("Failed. want= %v got= %v", test.want, got)
			}
		})
	}
}
