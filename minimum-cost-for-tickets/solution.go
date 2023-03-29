package minimumcostfortickets

import (
	"math"
	"sort"
)

// Tabular DP
func mincostTickets(days []int, costs []int) int {
	travelDays := make(map[int]struct{})
	for _, day := range days {
		travelDays[day] = struct{}{}
	}

	ticketCost := map[int]int{
		1:  costs[0],
		7:  costs[1],
		30: costs[2],
	}
	mincost := make([]int, 396)

	for i := days[len(days)-1]; i > 0; i-- {
		if _, found := travelDays[i]; !found {
			mincost[i] = mincost[i+1]
			continue
		}

		mincost[i] = math.MaxInt32
		for duration, cost := range ticketCost {
			mincost[i] = minOf(mincost[i], cost+mincost[i+duration])
		}
	}
	return mincost[1]
}

// Memoised
func mincostTickets3(days []int, costs []int) int {
	ticketCost := map[int]int{
		1:  costs[0],
		7:  costs[1],
		30: costs[2],
	}

	cache := map[int]int{len(days): 0}

	var mincost func(int) int
	mincost = func(i int) int {
		if cached, found := cache[i]; found {
			return cached
		}

		cache[i] = math.MaxInt32
		for duration, cost := range ticketCost {
			cache[i] = minOf(
				cache[i],
				cost+mincost(upperBound(days, days[i]+duration)),
			)
		}
		return cache[i]
	}

	return mincost(0)
}

// Recursion with closure
func mincostTickets2(days []int, costs []int) int {
	ticketCost := map[int]int{
		1:  costs[0],
		7:  costs[1],
		30: costs[2],
	}

	var mincost func(int) int
	mincost = func(i int) int {
		if i == len(days) {
			return 0
		}

		result := math.MaxInt32
		for duration, cost := range ticketCost {
			result = minOf(
				result,
				cost+mincost(upperBound(days, days[i]+duration)),
			)
		}
		return result
	}

	return mincost(0)
}

// Basic recursion
func mincostTickets1(days []int, costs []int) int {
	if len(days) == 0 {
		return 0
	}

	ticketCost := map[int]int{
		1:  costs[0],
		7:  costs[1],
		30: costs[2],
	}

	result := math.MaxInt32
	for duration, cost := range ticketCost {
		nextDayIndex := upperBound(days, days[0]+duration)
		result = minOf(result, cost+mincostTickets1(days[nextDayIndex:], costs))
	}

	return result
}

func upperBound(nums []int, n int) int {
	return sort.Search(len(nums), func(i int) bool {
		return nums[i] >= n
	})
}

func minOf(lhs, rhs int) int {
	if rhs < lhs {
		return rhs
	}
	return lhs
}
