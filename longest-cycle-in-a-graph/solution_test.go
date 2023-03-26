package longestcycleinagraph

import (
	"fmt"
	"reflect"
	"testing"
)

func TestLongestCycle(t *testing.T) {
	tests := []struct {
		edges []int
		want  int
	}{
		{[]int{3, 3, 4, 2, 3}, 3},
		{[]int{2, -1, 3, 1}, -1},
		{[]int{-1, 4, -1, 2, 0, 4}, -1},
	}

	for i, test := range tests {
		testName := fmt.Sprintf("Test_%v", i)
		t.Run(testName, func(t *testing.T) {

			got := longestCycle(test.edges)

			if !reflect.DeepEqual(got, test.want) {
				t.Errorf("Failed. want= %v got= %v", test.want, got)
			}
		})
	}
}
