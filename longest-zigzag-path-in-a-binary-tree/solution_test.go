package longestzigzagpathinabinarytree

import (
	"fmt"
	"reflect"
	"testing"
)

func TestLongestZigZag(t *testing.T) {
	tests := []struct {
		root *TreeNode
		want int
	}{
		{treeOf(1, nil, 1, 1, 1, nil, nil, 1, 1, nil, 1, nil, nil, nil, 1, nil, 1), 3},
		{treeOf(1, 1, 1, nil, 1, nil, nil, 1, 1, nil, 1), 4},
		{treeOf(1), 0},
	}

	for i, test := range tests {
		testName := fmt.Sprintf("Test_%v", i)
		t.Run(testName, func(t *testing.T) {
			got := longestZigZag(
				test.root,
			)

			if !reflect.DeepEqual(got, test.want) {
				t.Errorf("Failed. want= %v got= %v", test.want, got)
			}
		})
	}
}
