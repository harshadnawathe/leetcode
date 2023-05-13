package countwaystobuildgoodstrings

import (
	"fmt"
	"reflect"
	"testing"
)

func TestCountGoodStrings(t *testing.T) {
	tests := []struct {
		low  int
		high int
		zero int
		one  int
		want int
	}{
		{3, 3, 1, 1, 8},
		{2, 3, 1, 2, 5},
		{200, 200, 10, 1, 764262396},
	}

	for i, test := range tests {
		testName := fmt.Sprintf("Test_%v", i)
		t.Run(testName, func(t *testing.T) {
			got := countGoodStrings(
				test.low,
				test.high,
				test.zero,
				test.one,
			)

			if !reflect.DeepEqual(got, test.want) {
				t.Errorf("Failed. want= %v got= %v", test.want, got)
			}
		})
	}
}
