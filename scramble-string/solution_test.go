package scramblestring

import (
	"fmt"
	"reflect"
	"testing"
)

func TestIsScramble(t *testing.T) {

	tests := []struct {
		s1   string
		s2   string
		want bool
	}{
		{"great", "rgeat", true},
		{"abcde", "caebd", false},
		{"a", "a", true},
	}

	for i, test := range tests {
		testName := fmt.Sprintf("Test_%v", i)
		t.Run(testName, func(t *testing.T) {
			got := isScramble(
				test.s1,
				test.s2,
			)

			if !reflect.DeepEqual(got, test.want) {
				t.Errorf("Failed. want= %v got= %v", test.want, got)
			}
		})
	}
}
