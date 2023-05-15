package swappingnodesinalinkedlist

import (
	"fmt"
	"reflect"
	"testing"
)

func TestSwapNodes(t *testing.T) {
	tests := []struct {
		head *ListNode
		k    int
		want *ListNode
	}{
		{listOf(1, 2, 3, 4, 5), 2, listOf(1, 4, 3, 2, 5)},
		{listOf(7, 9, 6, 6, 7, 8, 3, 0, 9, 5), 5, listOf(7, 9, 6, 6, 8, 7, 3, 0, 9, 5)},
	}

	for i, test := range tests {
		testName := fmt.Sprintf("Test_%v", i)
		t.Run(testName, func(t *testing.T) {
			got := swapNodes(
				test.head,
				test.k,
			)

			if !reflect.DeepEqual(got, test.want) {
				t.Errorf("Failed. want= %v got= %v", test.want, got)
			}
		})
	}
}
