package stringtointegeratoi

import (
	"fmt"
	"reflect"
	"testing"
)

func TestMyAtoi(t *testing.T) {

	tests := []struct {
		s    string
		want int
	}{
		{"42", 42},
		{"   -42", -42},
		{"  4193", 4193},
		{"4192 some text", 4192},
		{"214748364712910", intMax},
		{"-214748364712910", intMin},
		{"2147483647", intMax},
		{"-2147483648", intMin},
	}

	for i, test := range tests {
		testName := fmt.Sprintf("Test_%v", i)
		t.Run(testName, func(t *testing.T) {
			got := myAtoi(
				test.s,
			)

			if !reflect.DeepEqual(got, test.want) {
				t.Errorf("Failed. want= %v got= %v", test.want, got)
			}
		})
	}
}
