package smallestevenmultiple

import (
	"fmt"
	"reflect"
	"testing"
)

func TestSmallestEvenMultiple(t *testing.T) {
	tests := []struct {
		n    int
		want int
	}{
		{1, 2},
		{2, 2},
		{5, 10},
		{6, 6},
	}

	for i, test := range tests {
		testName := fmt.Sprintf("Test_%v", i)
		t.Run(testName, func(t *testing.T) {
			got := smallestEvenMultiple(
				test.n,
			)

			if !reflect.DeepEqual(got, test.want) {
				t.Errorf("Failed. want= %v got= %v", test.want, got)
			}
		})
	}
}
