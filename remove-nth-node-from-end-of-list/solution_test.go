package removenthnodefromendoflist

import (
	"fmt"
	"testing"
)

func TestRemoveNthFromEnd(t *testing.T) {

	tests := []struct {
		head *ListNode
		n    int
		want *ListNode
	}{
		{listOf(1, 2, 3, 4, 5), 2, listOf(1, 2, 3, 5)},
		{listOf(1), 1, listOf()},
		{listOf(1, 2), 1, listOf(1)},
		{listOf(1, 2), 2, listOf(2)},
	}

	for i, test := range tests {
		testName := fmt.Sprintf("Test_%v", i)
		t.Run(testName, func(t *testing.T) {
			got := removeNthFromEnd(
				test.head,
				test.n,
			)

			if !equal(test.want, got) {
				t.Errorf("Failed. want= %v got= %v", test.want, got)
			}
		})
	}
}
