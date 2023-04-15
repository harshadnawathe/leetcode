package addstrings

import (
	"fmt"
	"reflect"
	"testing"
)

func TestAddStrings(t *testing.T) {
	tests := []struct {
		num1 string
		num2 string
		want string
	}{
		{"11", "123", "134"},
		{"99", "1", "100"},
		{"456", "77", "533"},
		{"0", "0", "0"},
	}

	for i, test := range tests {
		testName := fmt.Sprintf("Test_%v", i)
		t.Run(testName, func(t *testing.T) {
			got := addStrings(
				test.num1,
				test.num2,
			)

			if !reflect.DeepEqual(got, test.want) {
				t.Errorf("Failed. want= %v got= %v", test.want, got)
			}
		})
	}
}
