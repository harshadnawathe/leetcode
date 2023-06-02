package detonatethemaximumbombs

import (
	"fmt"
	"reflect"
	"testing"
)

func TestMaximumDetonation(t *testing.T) {
	tests := []struct {
		bombs [][]int
		want  int
	}{
		{[][]int{{2, 1, 3}, {6, 1, 4}}, 2},
		{[][]int{{1, 1, 5}, {10, 101, 5}}, 1},
		{[][]int{{1, 2, 3}, {2, 3, 1}, {3, 4, 2}, {4, 5, 3}, {5, 6, 4}}, 5},
		{[][]int{{4, 4, 3}, {4, 4, 3}}, 2},
	}

	for i, test := range tests {
		testName := fmt.Sprintf("Test_%v", i)
		t.Run(testName, func(t *testing.T) {
			got := maximumDetonation(
				test.bombs,
			)

			if !reflect.DeepEqual(got, test.want) {
				t.Errorf("Failed. want= %v got= %v", test.want, got)
			}
		})
	}
}
