package longestvalidparentheses

import (
	"fmt"
	"reflect"
	"testing"
)

func TestLongestValidParentheses(t *testing.T) {

	tests := []struct {
		s    string
		want int
	}{
		{"(()", 2},
		{"", 0},
		{")()())", 4},
		{"(()()", 4},
	}

	for i, test := range tests {
		testName := fmt.Sprintf("Test_%v", i)
		t.Run(testName, func(t *testing.T) {
			got := longestValidParentheses(
				test.s,
			)

			if !reflect.DeepEqual(got, test.want) {
				t.Errorf("Failed. want= %v got= %v", test.want, got)
			}
		})
	}
}
