package maximumwidthofbinarytree

import (
	"fmt"
	"reflect"
	"testing"
)

func TestWidthOfBinaryTree(t *testing.T) {
	tests := []struct {
		root *TreeNode
		want int
	}{
		{treeOf(1, 3, 2, 5, 3, nil, 9), 4},
		{treeOf(1, 3, 2, 5, nil, nil, 9, 6, nil, 7), 7},
		{treeOf(1, 3, 2, 5), 2},
	}

	for i, test := range tests {
		testName := fmt.Sprintf("Test_%v", i)
		t.Run(testName, func(t *testing.T) {
			got := widthOfBinaryTree(
				test.root,
			)

			if !reflect.DeepEqual(got, test.want) {
				t.Errorf("Failed. want= %v got= %v", test.want, got)
			}
		})
	}
}
