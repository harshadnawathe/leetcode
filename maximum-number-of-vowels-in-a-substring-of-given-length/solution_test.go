package maximumnumberofvowelsinasubstringofgivenlength

import (
	"fmt"
	"reflect"
	"testing"
)

func TestMaxVowels(t *testing.T) {
	tests := []struct {
		s    string
		k    int
		want int
	}{
		{"abciiidef", 3, 3},
		{"aeiou", 2, 2},
		{"leetcode", 3, 2},
		{"xyzlmn", 2, 0},
		{"abciou", 3, 3},
	}

	for i, test := range tests {
		testName := fmt.Sprintf("Test_%v", i)
		t.Run(testName, func(t *testing.T) {
			got := maxVowels(
				test.s,
				test.k,
			)

			if !reflect.DeepEqual(got, test.want) {
				t.Errorf("Failed. want= %v got= %v", test.want, got)
			}
		})
	}
}
