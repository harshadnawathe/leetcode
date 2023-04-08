package dividetwointegers

import (
	"fmt"
	"reflect"
	"testing"
)

func TestDivide(t *testing.T) {

	tests := []struct {
		dividend int
		divisor  int
		want     int
	}{
		{11, 2, 5},
		{7, -3, -2},
		{-1, -1, 1},
		{intMax + 1, 1, intMax},
		{intMax, -1, -intMax},
		{intMax + 1, -1, intMin},
	}

	for i, test := range tests {
		testName := fmt.Sprintf("Test_%v", i)
		t.Run(testName, func(t *testing.T) {
			got := divide(
				test.dividend,
				test.divisor,
			)

			if !reflect.DeepEqual(got, test.want) {
				t.Errorf("Failed. want= %v got= %v", test.want, got)
			}
		})
	}
}
