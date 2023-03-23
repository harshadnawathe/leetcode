package minimumscoreofapathbetweentwocities

import (
	"fmt"
	"reflect"
	"testing"
)

func TestMinScoore(t *testing.T) {
	tests := []struct {
		dest  int
		roads [][]int
		want  int
	}{
		{4, [][]int{{1, 2, 9}, {2, 3, 6}, {2, 4, 5}, {1, 4, 7}}, 5},
		{4, [][]int{{1, 2, 2}, {1, 3, 4}, {3, 4, 7}}, 2},
	}

	for i, test := range tests {
		testName := fmt.Sprintf("Test %d", i)
		t.Run(testName, func(t *testing.T) {
			got := minScore(test.dest, test.roads)

			if !reflect.DeepEqual(test.want, got) {
				t.Errorf("%s failed. want= %v got= %v", testName, test.want, got)
			}
		})
	}
}
