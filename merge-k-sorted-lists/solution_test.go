package mergeksortedlists

import (
	"fmt"
	"reflect"
	"testing"
)

func TestMergeKLists(t *testing.T) {

	tests := []struct {
		lists []*ListNode
		want  *ListNode
	}{
		{[]*ListNode{}, listOf()},
		{[]*ListNode{listOf()}, listOf()},
		{[]*ListNode{listOf(1, 4, 5), listOf(1, 3, 4), listOf(2, 6)}, listOf(1, 1, 2, 3, 4, 4, 5, 6)},
	}

	for i, test := range tests {
		testName := fmt.Sprintf("Test_%v", i)
		t.Run(testName, func(t *testing.T) {
			got := mergeKLists(
				test.lists,
			)

			if !reflect.DeepEqual(got, test.want) {
				t.Errorf("Failed. want= %v got= %v", test.want, got)
			}
		})
	}
}
