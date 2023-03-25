package romantointeger

import (
	"fmt"
	"reflect"
	"testing"
)

func TestRomanToInt(t *testing.T) {
	tests := []struct {
		s    string
		want int
	}{
		{"III", 3},
		{"LVIII", 58},
		{"MCMXCIV", 1994},
	}

	for i, test := range tests {
		testName := fmt.Sprintf("Test_%v", i)
		t.Run(testName, func(t *testing.T) {

			got := romanToInt(test.s)

			if !reflect.DeepEqual(got, test.want) {
				t.Errorf("Failed. want= %v got= %v", test.want, got)
			}
		})
	}
}
