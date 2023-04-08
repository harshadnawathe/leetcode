package clonegraph

import (
	"fmt"
	"testing"
)

func TestCloneGraph(t *testing.T) {

	tests := []struct {
		node *Node
		want *Node
	}{
		{nil, nil},
		{graphOf([][]int{{2, 4}, {1, 3}, {2, 4}, {1, 3}}), graphOf([][]int{{2, 4}, {1, 3}, {2, 4}, {1, 3}})},
		{graphOf([][]int{{}}), graphOf([][]int{{}})},
	}

	for i, test := range tests {
		testName := fmt.Sprintf("Test_%v", i)
		t.Run(testName, func(t *testing.T) {
			got := cloneGraph(
				test.node,
			)

			if !equal(got, test.want) {
				t.Errorf("Failed. want= %v got= %v", test.want, got)
			}
		})
	}
}
