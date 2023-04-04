package generateparentheses

import (
	"fmt"
	"testing"

	"github.com/google/go-cmp/cmp"
	"github.com/google/go-cmp/cmp/cmpopts"
)

var sortAscending = cmpopts.SortSlices(func(a, b string) bool {
	return a < b
})

func TestGenerateParenthesis(t *testing.T) {

	tests := []struct {
		n    int
		want []string
	}{
		{0, []string{""}},
		{1, []string{"()"}},
		{2, []string{"()()", "(())"}},
		// {3, []string{"((()))", "(()())", "(())()", "()(())", "()()()"}},
		// {4, []string{"(((())))", "((()()))", "((())())", "((()))()", "(()(()))", "(()()())", "(()())()", "(())(())", "(())()()", "()((()))", "()(()())", "()(())()", "()()(())", "()()()()"}},
	}

	for i, test := range tests {
		testName := fmt.Sprintf("Test_%v", i)
		t.Run(testName, func(t *testing.T) {
			got := generateParenthesis(
				test.n,
			)

			if !cmp.Equal(got, test.want, sortAscending) {
				t.Errorf("Failed. want= %v got= %v", test.want, got)
			}
		})
	}
}
