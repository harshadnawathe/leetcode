package longestpalindromicsubsequence

import (
	"fmt"
	"reflect"
	"testing"
)

func TestLongestPalindromeSubseq(t *testing.T) {

	tests := []struct {
		s    string
		want int
	}{
		{"", 0},
		{"a", 1},
		{"bb", 2},
		{"bbbab", 4},
		{"cbbd", 2},
	}

	for i, test := range tests {
		testName := fmt.Sprintf("Test_%v", i)
		t.Run(testName, func(t *testing.T) {
			got := longestPalindromeSubseq(
				test.s,
			)

			if !reflect.DeepEqual(got, test.want) {
				t.Errorf("Failed. want= %v got= %v", test.want, got)
			}
		})
	}
}
