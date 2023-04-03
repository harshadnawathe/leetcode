package mergetwosortedlists

import (
	"fmt"
	"reflect"
	"testing"
)

func TestMergeTwoLists(t *testing.T) {

	tests := []struct {
		list1 *ListNode
		list2 *ListNode
		want  *ListNode
	}{
		{nil, nil, nil},
		{listOf(1), nil, listOf(1)},
		{nil, listOf(1), listOf(1)},
		{listOf(1), listOf(1), listOf(1, 1)},
		{listOf(2), listOf(1), listOf(1, 2)},
		{listOf(1, 2, 4), listOf(1, 3, 4), listOf(1, 1, 2, 3, 4, 4)},
	}

	for i, test := range tests {
		testName := fmt.Sprintf("Test_%v", i)
		t.Run(testName, func(t *testing.T) {
			got := mergeTwoLists(
				test.list1,
				test.list2,
			)

			if !reflect.DeepEqual(got, test.want) {
				t.Errorf("Failed. want= %v got= %v", test.want, got)
			}
		})
	}
}
