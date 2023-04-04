package generateparentheses

import "fmt"

func generateParenthesis1(n int) []string {
	if n == 0 {
		return []string{""}
	}

	var result []string
	for i := 0; i < n; i++ {
		for _, left := range generateParenthesis1(i) {
			for _, right := range generateParenthesis1(n - i - 1) {
				result = append(result, fmt.Sprintf("(%s)%s", left, right))
			}
		}
	}
	return result
}

// Backtracking
func generateParenthesis(n int) []string {
	var result []string

	var f func(string, int, int)
	f = func(s string, nOpen, nClose int) {
		if len(s) == 2*n {
			result = append(result, s)
			return
		}

		if nOpen < n {
			f(s+"(", nOpen+1, nClose)
		}

		if nClose < nOpen {
			f(s+")", nOpen, nClose+1)
		}
	}
	f("", 0, 0)
	return result
}
