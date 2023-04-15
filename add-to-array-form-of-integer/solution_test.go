package addtoarrayformofinteger

import (
	"fmt"
	"reflect"
	"testing"
)

func TestAddToArrayForm(t *testing.T) {
	tests := []struct {
		num  []int
		k    int
		want []int
	}{
		{[]int{1, 0}, 10, []int{2, 0}},
		{[]int{9, 9}, 1, []int{1, 0, 0}},
		{[]int{1, 2, 0, 0}, 34, []int{1, 2, 3, 4}},
		{[]int{2, 7, 4}, 181, []int{4, 5, 5}},
		{[]int{2, 1, 5}, 806, []int{1, 0, 2, 1}},
	}

	for i, test := range tests {
		testName := fmt.Sprintf("Test_%v", i)
		t.Run(testName, func(t *testing.T) {
			got := addToArrayForm(
				test.num,
				test.k,
			)

			if !reflect.DeepEqual(got, test.want) {
				t.Errorf("Failed. want= %v got= %v", test.want, got)
			}
		})
	}
}
