package longestsubstringwithoutrepeatingcharacters

import (
	"fmt"
	"testing"
)

func TestLengthOfLongestSubstring(t *testing.T) {
	tests := []struct {
		s    string
		want int
	}{
		{"abcabcbb", 3},
		{"bbbbb", 1},
		{"pwwkew", 3},
	}

	for i, test := range tests {
		testName := fmt.Sprintf("Test_%d", i)
		t.Run(testName, func(t *testing.T) {
			got := lengthOfLongestSubstring(test.s)

			if got != test.want {
				t.Errorf("Failed. want= %v got= %v", test.want, got)
			}
		})
	}
}
