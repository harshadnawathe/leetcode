package russiandollenvelopes

import (
	"fmt"
	"reflect"
	"testing"
)

func TestMaxEnvelopes(t *testing.T) {
	tests := []struct {
		envelopes [][]int
		want      int
	}{
		{[][]int{{5, 4}, {6, 4}, {6, 7}, {2, 3}}, 3},
		{[][]int{{1, 1}, {1, 1}, {1, 1}}, 1},
	}

	for i, test := range tests {
		testName := fmt.Sprintf("Test_%v", i)
		t.Run(testName, func(t *testing.T) {
			got := maxEnvelopes(
				test.envelopes,
			)

			if !reflect.DeepEqual(got, test.want) {
				t.Errorf("Failed. want= %v got= %v", test.want, got)
			}
		})
	}
}
