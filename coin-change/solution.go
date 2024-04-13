package coinchange

const infinity = int(^uint(0) >> 1)

func coinChange(coins []int, amount int) int {
	minCoins := make([]int, amount+1)
	for amt := 1; amt <= amount; amt++ {
		minCoins[amt] = infinity
		for _, coin := range coins {
			if coin <= amt && minCoins[amt-coin] != infinity {
				minCoins[amt] = min(minCoins[amt], minCoins[amt-coin]+1)
			}
		}
	}

	if minCoins[amount] == infinity {
		return -1
	}

	return minCoins[amount]
}
