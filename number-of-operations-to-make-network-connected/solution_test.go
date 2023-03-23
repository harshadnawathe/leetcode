package numberofoperationstomakenetworkconnected

import (
	"fmt"
	"reflect"
	"testing"
)

func TestMakeConnected(t *testing.T) {
	tests := []struct {
		n           int
		connections [][]int
		want        int
	}{
		{4, [][]int{{0, 1}, {0, 2}, {1, 2}}, 1},
		{6, [][]int{{0, 1}, {0, 2}, {0, 3}, {1, 2}, {1, 3}}, 2},
		{6, [][]int{{0, 1}, {0, 2}, {0, 3}, {1, 2}}, -1},
	}

	for i, test := range tests {
		testName := fmt.Sprintf("Test %d", i)
		t.Run(testName, func(t *testing.T) {
			got := makeConnected(test.n, test.connections)

			if !reflect.DeepEqual(got, test.want) {
				t.Errorf("Failed. want= %v got= %v", test.want, got)
			}

		})
	}

}
