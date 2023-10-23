package poweroffour

import (
	"fmt"
	"reflect"
	"testing"
)

func TestIsPowerOfFour(t *testing.T) {
	tests := []struct {
		n    int
		want bool
	}{{1, true}, {2, false}, {4, true}, {-16, false}, {256, true}}

	for i, test := range tests {
		testName := fmt.Sprintf("Test_%v", i)
		t.Run(testName, func(t *testing.T) {
			got := isPowerOfFour(
				test.n,
			)

			if !reflect.DeepEqual(got, test.want) {
				t.Errorf("Failed. want= %v got= %v", test.want, got)
			}
		})
	}
}
