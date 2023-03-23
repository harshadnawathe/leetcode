package longestpalindromicsubstring

import (
	"fmt"
	"testing"
)

func TestLongestPalindrome(t *testing.T) {
	tests := []struct {
		s    string
		want string
	}{
		{"babad", "bab"},
		{"cbbd", "bb"},
	}

	for i, test := range tests {
		testName := fmt.Sprintf("Test_%v", i)
		t.Run(testName, func(t *testing.T) {
			got := longestPalindrome(test.s)

			if got != test.want {
				t.Errorf("Failed. got= %v want= %v", got, test.want)
			}
		})
	}

}
