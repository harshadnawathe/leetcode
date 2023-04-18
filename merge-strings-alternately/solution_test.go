package mergestringsalternately

import (
	"fmt"
	"reflect"
	"testing"
)

func TestMergeAlternately(t *testing.T) {
	tests := []struct {
		word1 string
		word2 string
		want  string
	}{
		{"abc", "pqr", "apbqcr"},
		{"ab", "pqrs", "apbqrs"},
		{"abcd", "pq", "apbqcd"},
	}

	for i, test := range tests {
		testName := fmt.Sprintf("Test_%v", i)
		t.Run(testName, func(t *testing.T) {
			got := mergeAlternately(
				test.word1,
				test.word2,
			)

			if !reflect.DeepEqual(got, test.want) {
				t.Errorf("Failed. want= %v got= %v", test.want, got)
			}
		})
	}
}
