package swapnodesinpairs

import (
	"fmt"
	"testing"
)

func TestSwapPairs(t *testing.T) {

	tests := []struct {
		head *ListNode
		want *ListNode
	}{
		{listOf(1, 2, 3, 4), listOf(2, 1, 4, 3)},
		{listOf(), listOf()},
		{listOf(1), listOf(1)},
		{listOf(1, 2, 3), listOf(2, 1, 3)},
	}

	for i, test := range tests {
		testName := fmt.Sprintf("Test_%v", i)
		t.Run(testName, func(t *testing.T) {
			got := swapPairs(
				test.head,
			)

			if !equal(got, test.want) {
				t.Errorf("Failed. want= %v got= %v", test.want, got)
			}
		})
	}
}
