package relativeranks

import (
	"sort"
	"strconv"
)

func findRelativeRanks(score []int) []string {
	players := make([]int, len(score))
	for i := range players {
		players[i] = i
	}

	sort.Slice(players, func(i, j int) bool {
		return score[players[i]] > score[players[j]]
	})

	ranks := make([]string, len(score))

	medal := []string{"Gold Medal", "Silver Medal", "Bronze Medal"}

	for rank, player := range players {
		if rank < len(medal) {
			ranks[player] = medal[rank]
		} else {
			ranks[player] = strconv.Itoa(rank + 1)
		}
	}

	return ranks
}
