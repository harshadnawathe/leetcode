package restorethearray

import (
	"fmt"
	"reflect"
	"testing"
)

func TestNumberOfArrays(t *testing.T) {
	tests := []struct {
		s    string
		k    int
		want int
	}{
		{"1000", 10000, 1},
		{"1000", 10, 0},
		{"1317", 2000, 8},
		{"2020", 30, 1},
	}

	for i, test := range tests {
		testName := fmt.Sprintf("Test_%v", i)
		t.Run(testName, func(t *testing.T) {
			got := numberOfArrays(
				test.s,
				test.k,
			)

			if !reflect.DeepEqual(got, test.want) {
				t.Errorf("Failed. want= %v got= %v", test.want, got)
			}
		})
	}
}
