package bulbswitcher

import (
	"fmt"
	"reflect"
	"testing"
)

func TestBulbSwitch(t *testing.T) {
	tests := []struct {
		n    int
		want int
	}{
		{10, 3},
		{50, 7},
		{3, 1},
		{0, 0},
		{1, 1},
	}

	for i, test := range tests {
		testName := fmt.Sprintf("Test_%v", i)
		t.Run(testName, func(t *testing.T) {
			got := bulbSwitch(
				test.n,
			)

			if !reflect.DeepEqual(got, test.want) {
				t.Errorf("Failed. want= %v got= %v", test.want, got)
			}
		})
	}
}
