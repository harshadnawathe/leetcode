package validparentheses

import (
	"fmt"
	"reflect"
	"testing"
)

func TestIsValid(t *testing.T) {

	tests := []struct {
		s    string
		want bool
	}{
		{"", true},
		{"()", true},
		{"(]", false},
		{"{()[]", false},
		{"{}[]()", true},
		{"]", false},
	}

	for i, test := range tests {
		testName := fmt.Sprintf("Test_%v", i)
		t.Run(testName, func(t *testing.T) {
			got := isValid(
				test.s,
			)

			if !reflect.DeepEqual(got, test.want) {
				t.Errorf("Failed. want= %v got= %v", test.want, got)
			}
		})
	}
}
