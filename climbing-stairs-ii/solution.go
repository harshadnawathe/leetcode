// Package climbingstairsii solves https://leetcode.com/problems/climbing-stairs-ii/
package climbingstairsii

// time: O(n), space: O(n)
func climbStairs(n int, costs []int) int {
	cost := func(i, j int) int {
		return costs[j-1] + (j-i)*(j-i)
	}

	stepMinCost := make([]int, n+1)
	stepMinCost[0] = 0
	if len(stepMinCost) > 1 {
		stepMinCost[1] = cost(0, 1)
	}
	if len(stepMinCost) > 2 {
		stepMinCost[2] = min(stepMinCost[1]+cost(1, 2), cost(0, 2))
	}

	for step := 3; step <= n; step++ {
		stepMinCost[step] = min(
			stepMinCost[step-3]+cost(step-3, step),
			stepMinCost[step-2]+cost(step-2, step),
			stepMinCost[step-1]+cost(step-1, step),
		)
	}

	return stepMinCost[n]
}
