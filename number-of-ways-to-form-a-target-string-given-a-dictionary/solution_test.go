package numberofwaystoformatargetstringgivenadictionary

import (
	"fmt"
	"reflect"
	"testing"
)

func TestNumWays(t *testing.T) {
	tests := []struct {
		words  []string
		target string
		want   int
	}{
		{[]string{"acca", "bbbb", "caca"}, "aba", 6},
		{[]string{"abba", "baab"}, "bab", 4},
	}

	for i, test := range tests {
		testName := fmt.Sprintf("Test_%v", i)
		t.Run(testName, func(t *testing.T) {
			got := numWays(
				test.words,
				test.target,
			)

			if !reflect.DeepEqual(got, test.want) {
				t.Errorf("Failed. want= %v got= %v", test.want, got)
			}
		})
	}
}
