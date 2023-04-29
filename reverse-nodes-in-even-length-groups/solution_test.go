package reversenodesinevenlengthgroups

import (
	"fmt"
	"reflect"
	"testing"
)

func TestReverseEvenLengthGroups(t *testing.T) {
	tests := []struct {
		head *ListNode
		want *ListNode
	}{
		{listOf(5, 2, 6, 3, 9, 1, 7, 3, 8, 4), listOf(5, 6, 2, 3, 9, 1, 4, 8, 3, 7)},
		{listOf(1, 1, 0, 6), listOf(1, 0, 1, 6)},
		{listOf(1, 1, 0, 6, 5), listOf(1, 0, 1, 5, 6)},
		{listOf(1, 2, 3), listOf(1, 3, 2)},
	}

	for i, test := range tests {
		testName := fmt.Sprintf("Test_%v", i)
		t.Run(testName, func(t *testing.T) {
			got := reverseEvenLengthGroups(
				test.head,
			)

			if !reflect.DeepEqual(got, test.want) {
				t.Errorf("Failed. want= %v got= %v", test.want, got)
			}
		})
	}
}
