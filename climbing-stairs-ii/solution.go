// Package climbingstairsii solves https://leetcode.com/problems/climbing-stairs-ii/
package climbingstairsii

// time: O(n), space: O(1)
func climbStairs(n int, costs []int) int {
	a, b, c := 0, 0, 0

	for i := range n {
		next := min(min(a+9, b+4), c+1) + costs[i]
		a, b, c = b, c, next
	}

	return c
}
