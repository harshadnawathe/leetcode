package openthelock

func openLock(deadends []string, target string) int {
	nextSlot := map[byte][]string{
		'0': {"1", "9"},
		'1': {"2", "0"},
		'2': {"3", "1"},
		'3': {"4", "2"},
		'4': {"5", "3"},
		'5': {"6", "4"},
		'6': {"7", "5"},
		'7': {"8", "6"},
		'8': {"9", "7"},
		'9': {"0", "8"},
	}

	deadendSet := make(map[string]struct{})
	for _, deadend := range deadends {
		deadendSet[deadend] = struct{}{}
	}

	type Lock struct {
		Combination string
		NumTurns    int
	}

	visited := make(map[string]struct{})

	q := []Lock{{"0000", 0}}

	for len(q) > 0 {
		lock := q[0]
		q = q[1:]

		if lock.Combination == target {
			return lock.NumTurns
		}

		if _, ok := deadendSet[lock.Combination]; ok {
			continue
		}

		if _, ok := visited[lock.Combination]; ok {
			continue
		}

		visited[lock.Combination] = struct{}{}

		for i := 0; i < len(lock.Combination); i++ {
			for _, next := range nextSlot[lock.Combination[i]] {
				nextCombination := lock.Combination[:i] + next + lock.Combination[i+1:]

				if _, ok := deadendSet[nextCombination]; !ok {
					q = append(q,
						Lock{
							nextCombination,
							lock.NumTurns + 1,
						})
				}

			}
		}

	}

	return -1
}
