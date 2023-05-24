package wordsearchii

import (
	"fmt"
	"testing"

	"github.com/google/go-cmp/cmp"
	"github.com/google/go-cmp/cmp/cmpopts"
)

var sortAscending = cmpopts.SortSlices(func(a, b string) bool {
	return a < b
})

func TestFindWords(t *testing.T) {
	tests := []struct {
		board [][]byte
		words []string
		want  []string
	}{
		{[][]byte{{'o', 'a', 'a', 'n'}, {'e', 't', 'a', 'e'}, {'i', 'h', 'k', 'r'}, {'i', 'f', 'l', 'v'}}, []string{"oath", "pea", "eat", "rain"}, []string{"eat", "oath"}},
		{[][]byte{{'a', 'b'}, {'c', 'd'}}, []string{"abcb"}, []string{}},
	}

	for i, test := range tests {
		testName := fmt.Sprintf("Test_%v", i)
		t.Run(testName, func(t *testing.T) {
			got := findWords(
				test.board,
				test.words,
			)

			if !cmp.Equal(got, test.want, sortAscending) {
				t.Errorf("Failed. want= %v got= %v", test.want, got)
			}
		})
	}
}
