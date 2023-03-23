package addtwonumbers

import (
	"fmt"
	"testing"
)

func TestAddTwoNumbers(t *testing.T) {
	tests := []struct {
		l1, l2 *ListNode
		want   *ListNode
	}{
		{
			makeLinkedList(2, 4, 3),
			makeLinkedList(5, 6, 4),
			makeLinkedList(7, 0, 8),
		},
		{
			makeLinkedList(0),
			makeLinkedList(0),
			makeLinkedList(0),
		},
		{
			makeLinkedList(9, 9, 9, 9, 9, 9, 9),
			makeLinkedList(9, 9, 9, 9),
			makeLinkedList(8, 9, 9, 9, 0, 0, 0, 1),
		},
	}

	for i, test := range tests {
		testName := fmt.Sprintf("Test_%v", i)
		t.Run(testName, func(t *testing.T) {
			got := addTwoNumbers(test.l1, test.l2)
			if !equal(got, test.want) {
				t.Errorf("Failed. want= %v got= %v", test.want, got)
			}
		})
	}

}
