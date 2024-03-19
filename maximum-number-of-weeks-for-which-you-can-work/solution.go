package maximumnumberofweeksforwhichyoucanwork

// time: O(n)
// space: O(1)
func numberOfWeeks(milestones []int) int64 {
	totalWeeks, maxMilestone := 0, 0
	for _, milestone := range milestones {
		totalWeeks += milestone

		if maxMilestone < milestone {
			maxMilestone = milestone
		}
	}

	remainingMilestones := totalWeeks - maxMilestone
	if maxMilestone > remainingMilestones+1 {
		maxMilestone = remainingMilestones + 1
	}

	return int64(maxMilestone + remainingMilestones)
}
