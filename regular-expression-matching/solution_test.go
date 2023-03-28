package regularexpressionmatching

import (
	"fmt"
	"reflect"
	"testing"
)

func TestIsMatch(t *testing.T) {

	tests := []struct {
		s    string
		p    string
		want bool
	}{
		{"", "", true},
		{"", ".*", true},
		{"a", "a", true},
		{"a", "", false},
		{"", "a", false},
		{"aa", "a", false},
		{"abc", "a.c", true},
		{"aa", "a", false},
		{"aa", "a*", true},
		{"ab", ".*", true},
		{"ac", "ab*c", true},
		{"abc", "ab*c", true},
		{"abbc", "ab*c", true},
	}

	for i, test := range tests {
		testName := fmt.Sprintf("Test_%v", i)
		t.Run(testName, func(t *testing.T) {
			got := isMatch(
				test.s,
				test.p,
			)

			if !reflect.DeepEqual(got, test.want) {
				t.Errorf("Failed. want= %v got= %v", test.want, got)
			}
		})
	}
}
