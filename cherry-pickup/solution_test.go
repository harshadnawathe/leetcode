package cherrypickup

import (
	"fmt"
	"reflect"
	"testing"
)

func TestCherryPickup(t *testing.T) {
	tests := []struct {
		grid [][]int
		want int
	}{
		{[][]int{{0, 1, -1}, {1, 0, -1}, {1, 1, 1}}, 5},
		{[][]int{{1, 1, -1}, {1, -1, 1}, {-1, 1, 1}}, 0},
		{[][]int{{1, 1, -1, 1, 1}, {1, 1, 1, 1, 1}, {-1, 1, 1, -1, -1}, {0, 1, 1, -1, 0}, {1, 0, -1, 1, 0}}, 0},
	}

	for i, test := range tests {
		testName := fmt.Sprintf("Test_%v", i)
		t.Run(testName, func(t *testing.T) {
			got := cherryPickup(
				test.grid,
			)

			if !reflect.DeepEqual(got, test.want) {
				t.Errorf("Failed. want= %v got= %v", test.want, got)
			}
		})
	}
}
