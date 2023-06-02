package detonatethemaximumbombs

import (
	"math"
)

func maximumDetonation(bombs [][]int) int {
	adjList := makeAdjacencyList(bombs)

	var numDetonations func(int, map[int]struct{}) int
	numDetonations = func(bomb int, visited map[int]struct{}) int {
		visited[bomb] = struct{}{}
		num := 1
		for _, next := range adjList[bomb] {
			if _, found := visited[next]; found {
				continue
			}
			num += numDetonations(next, visited)
		}
		return num
	}

	max := 0
	for bomb := range bombs {
		if num := numDetonations(bomb, make(map[int]struct{})); max < num {
			max = num
		}
	}
	return max
}

func distance(x1, y1 int, x2, y2 int) float64 {
	return math.Sqrt(math.Pow(float64(x1-x2), 2) + math.Pow(float64(y1-y2), 2))
}

func makeAdjacencyList(bombs [][]int) map[int][]int {
	adjList := make(map[int][]int)

	for i, bomb1 := range bombs {
		for j := i + 1; j < len(bombs); j++ {
			bomb2 := bombs[j]
			dist := distance(bomb1[0], bomb1[1], bomb2[0], bomb2[1])

			if dist <= float64(bomb1[2]) {
				adjList[i] = append(adjList[i], j)
			}

			if dist <= float64(bomb2[2]) {
				adjList[j] = append(adjList[j], i)
			}
		}
	}

	return adjList
}
