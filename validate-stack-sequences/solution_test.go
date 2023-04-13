package validatestacksequences

import (
	"fmt"
	"reflect"
	"testing"
)

func TestValidateStackSequences(t *testing.T) {
	tests := []struct {
		pushed []int
		popped []int
		want   bool
	}{
		{[]int{1, 2, 3, 4, 5}, []int{4, 5, 3, 2, 1}, true},
		{[]int{1, 2, 3, 4, 5}, []int{4, 3, 5, 1, 2}, false},
		{[]int{1, 0}, []int{1, 0}, true},
	}

	for i, test := range tests {
		testName := fmt.Sprintf("Test_%v", i)
		t.Run(testName, func(t *testing.T) {
			got := validateStackSequences(
				test.pushed,
				test.popped,
			)

			if !reflect.DeepEqual(got, test.want) {
				t.Errorf("Failed. want= %v got= %v", test.want, got)
			}
		})
	}
}
