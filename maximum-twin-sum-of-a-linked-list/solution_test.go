package maximumtwinsumofalinkedlist

import (
	"fmt"
	"reflect"
	"testing"
)

func TestPairSum(t *testing.T) {
	tests := []struct {
		head *ListNode
		want int
	}{
		{listOf(5, 4, 2, 1), 6},
		{listOf(4, 2, 2, 3), 7},
		{listOf(1, 100000), 100001},
	}

	for i, test := range tests {
		testName := fmt.Sprintf("Test_%v", i)
		t.Run(testName, func(t *testing.T) {
			got := pairSum(
				test.head,
			)

			if !reflect.DeepEqual(got, test.want) {
				t.Errorf("Failed. want= %v got= %v", test.want, got)
			}
		})
	}
}
