package findplayerswithzerooronelosses

import (
	"sort"
)

func findWinners(matches [][]int) [][]int {
	playerLosses := map[int]int{}
	for _, match := range matches {
		winner, loser := match[0], match[1]
		playerLosses[loser]++
		if _, ok := playerLosses[winner]; !ok {
			playerLosses[winner] = 0
		}
	}

	var players []int
	for player := range playerLosses {
		players = append(players, player)
	}

	sort.Ints(players)

	neverLost := []int{}
	lostOnes := []int{}
	for _, player := range players {
		losses := playerLosses[player]
		if losses == 0 {
			neverLost = append(neverLost, player)
		} else if losses == 1 {
			lostOnes = append(lostOnes, player)
		}
	}

	return [][]int{neverLost, lostOnes}
}
