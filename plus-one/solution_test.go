package plusone

import (
	"fmt"
	"reflect"
	"testing"
)

func TestPlusOne(t *testing.T) {

	tests := []struct {
		digits []int
		want   []int
	}{
		{[]int{9}, []int{1, 0}},
		{[]int{1}, []int{2}},
		{[]int{1, 1, 9}, []int{1, 2, 0}},
		{[]int{1, 2, 3}, []int{1, 2, 4}},
		{[]int{4, 3, 2, 1}, []int{4, 3, 2, 2}},
	}

	for i, test := range tests {
		testName := fmt.Sprintf("Test_%v", i)
		t.Run(testName, func(t *testing.T) {
			got := plusOne(
				test.digits,
			)

			if !reflect.DeepEqual(got, test.want) {
				t.Errorf("Failed. want= %v got= %v", test.want, got)
			}
		})
	}
}
