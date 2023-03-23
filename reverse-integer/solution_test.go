package reverseinteger

import (
	"fmt"
	"reflect"
	"testing"
)

func TestReverse(t *testing.T) {
	tests := []struct {
		x    int
		want int
	}{
		{123, 321},
		{-123, -321},
		{120, 21},
		{1534236469, 0},
	}

	for i, test := range tests {
		testName := fmt.Sprintf("Test_%v", i)
		t.Run(testName, func(t *testing.T) {

			got := reverse(test.x)

			if !reflect.DeepEqual(got, test.want) {
				t.Errorf("Failed. want= %v got= %v", test.want, got)
			}
		})
	}
}
