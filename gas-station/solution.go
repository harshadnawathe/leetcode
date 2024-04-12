package gasstation

// time: O(n)
// space: O(1)
func canCompleteCircuit(gas []int, cost []int) int {
	start := 0
	tank := 0
	total := 0

	for i := range gas {
		tank += gas[i] - cost[i]
		total += gas[i] - cost[i]

		if tank < 0 {
			start = i + 1
			tank = 0
		}
	}

	if total < 0 {
		return -1
	}

	return start
}
