package numberofwaystostayinthesameplaceaftersomesteps

import (
	"fmt"
	"reflect"
	"testing"
)

func TestNumWays(t *testing.T) {
	tests := []struct {
		steps  int
		arrLen int
		want   int
	}{
		{
			steps:  3,
			arrLen: 2,
			want:   4,
		},
		{
			steps:  2,
			arrLen: 4,
			want:   2,
		},
		{
			steps:  4,
			arrLen: 2,
			want:   8,
		},
		{
			steps:  27,
			arrLen: 7,
			want:   127784505,
		},
	}

	for i, test := range tests {
		testName := fmt.Sprintf("Test_%v", i)
		t.Run(testName, func(t *testing.T) {
			got := numWays(
				test.steps,
				test.arrLen,
			)

			if !reflect.DeepEqual(got, test.want) {
				t.Errorf("Failed. want= %v got= %v", test.want, got)
			}
		})
	}
}
