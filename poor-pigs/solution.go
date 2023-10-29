package poorpigs

import "math"

func poorPigs(buckets int, minutesToDie int, minutesToTest int) int {
	maxNumTests := minutesToTest / minutesToDie
	numPigs := 0

	base := float64(maxNumTests + 1)

	for int(math.Pow(base, float64(numPigs))) < buckets {
		numPigs++
	}
	return numPigs
}
