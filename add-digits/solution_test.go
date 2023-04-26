package adddigits

import (
	"fmt"
	"reflect"
	"testing"
)

func TestAddDigits(t *testing.T) {
	tests := []struct {
		num  int
		want int
	}{
		{38, 2},
		{0, 0},
		{9, 9},
		{18, 9},
	}

	for i, test := range tests {
		testName := fmt.Sprintf("Test_%v", i)
		t.Run(testName, func(t *testing.T) {
			got := addDigits(
				test.num,
			)

			if !reflect.DeepEqual(got, test.want) {
				t.Errorf("Failed. want= %v got= %v", test.want, got)
			}
		})
	}
}
