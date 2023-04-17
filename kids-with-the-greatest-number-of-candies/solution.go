package kidswiththegreatestnumberofcandies

import "math"

func kidsWithCandies(candies []int, extraCandies int) []bool {
	max := math.MinInt32
	for _, n := range candies {
		if max < n {
			max = n
		}
	}

	result := make([]bool, len(candies))
	for i := range candies {
		result[i] = candies[i]+extraCandies >= max
	}

	return result
}
