package reversenodesinkgroup

import (
	"fmt"
	"testing"
)

func TestReverseKGroup(t *testing.T) {

	tests := []struct {
		head *ListNode
		k    int
		want *ListNode
	}{
		{nil, 0, nil},
		{listOf(1, 2), 1, listOf(1, 2)},
		{listOf(1, 2, 3, 4), 2, listOf(2, 1, 4, 3)},
		{listOf(1, 2, 3, 4), 3, listOf(3, 2, 1, 4)},
	}

	for i, test := range tests {
		testName := fmt.Sprintf("Test_%v", i)
		t.Run(testName, func(t *testing.T) {
			got := reverseKGroup(
				test.head,
				test.k,
			)

			if !equal(got, test.want) {
				t.Errorf("Failed. want= %v got= %v", test.want, got)
			}
		})
	}
}
