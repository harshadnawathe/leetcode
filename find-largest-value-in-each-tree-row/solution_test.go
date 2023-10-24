package findlargestvalueineachtreerow

import (
	"fmt"
	"reflect"
	"testing"
)

func TestLargestValues(t *testing.T) {
	tests := []struct {
		root *TreeNode
		want []int
	}{
		{treeOf(1, 3, 2, 5, 3, nil, 9), []int{1, 3, 9}},
		{treeOf(1, 2, 3), []int{1, 3}},
	}

	for i, test := range tests {
		testName := fmt.Sprintf("Test_%v", i)
		t.Run(testName, func(t *testing.T) {
			got := largestValues(
				test.root,
			)

			if !reflect.DeepEqual(got, test.want) {
				t.Errorf("Failed. want= %v got= %v", test.want, got)
			}
		})
	}
}
