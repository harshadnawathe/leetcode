package numberofstudentsunabletoeatlunch

func countStudents(students []int, sandwiches []int) int {
	count := [2]int{0, 0}

	for _, student := range students {
		count[student]++
	}

	for _, sandwich := range sandwiches {
		if count[sandwich] == 0 {
			return count[0] + count[1]
		}
		count[sandwich]--
	}

	return 0
}
