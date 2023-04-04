package optimalpartitionofstring

import (
	"fmt"
	"reflect"
	"testing"
)

func TestPartitionString(t *testing.T) {

	tests := []struct {
		s    string
		want int
	}{
		{"", 0},
		{"a", 1},
		{"ab", 1},
		{"aa", 2},
		{"abacaba", 4},
		{"ssssss", 6},
	}

	for i, test := range tests {
		testName := fmt.Sprintf("Test_%v", i)
		t.Run(testName, func(t *testing.T) {
			got := partitionString(
				test.s,
			)

			if !reflect.DeepEqual(got, test.want) {
				t.Errorf("Failed. want= %v got= %v", test.want, got)
			}
		})
	}
}
