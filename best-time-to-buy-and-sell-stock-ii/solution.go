package besttimetobuyandsellstockii

// time: O(n)
// space: O(1)
func maxProfit(prices []int) int {
	profitTotal := 0
	buyPrice := prices[0]

	for i := 1; i < len(prices); i++ {
		if prices[i-1] > prices[i] {
			profitTotal += prices[i-1] - buyPrice
			buyPrice = prices[i]
		}
	}

	return profitTotal + prices[len(prices)-1] - buyPrice
}
