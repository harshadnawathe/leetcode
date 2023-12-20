package buytwochocolates

import "sort"

func buyChoco(prices []int, money int) int {
	sort.Ints(prices)
	price := prices[0] + prices[1]
	if price > money {
		return money
	}
	return money - price
}
