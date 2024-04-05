package besttimetobuyandsellstock

// time: O(n)
// space: O(1)
func maxProfit(prices []int) int {
	profit := 0
	minIdx, maxIdx := 0, 0

	for i := 1; i < len(prices); i++ {
		price := prices[i]

		if price < prices[minIdx] {
			minIdx = i
			maxIdx = i
		} else if price > prices[maxIdx] {
			maxIdx = i
		}

		profit = max(profit, prices[maxIdx]-prices[minIdx])
	}

	return profit
}
