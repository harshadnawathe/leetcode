package laststoneweight

import (
	"fmt"
	"reflect"
	"testing"
)

func TestLastStoneWeight(t *testing.T) {
	tests := []struct {
		stones []int
		want   int
	}{
		{[]int{}, 0},
		{[]int{1}, 1},
		{[]int{2, 7, 4, 1, 8, 1}, 1},
		{[]int{2, 2}, 0},
	}

	for i, test := range tests {
		testName := fmt.Sprintf("Test_%v", i)
		t.Run(testName, func(t *testing.T) {
			got := lastStoneWeight(
				test.stones,
			)

			if !reflect.DeepEqual(got, test.want) {
				t.Errorf("Failed. want= %v got= %v", test.want, got)
			}
		})
	}
}
