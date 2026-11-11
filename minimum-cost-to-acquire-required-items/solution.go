package minimumcosttoacquirerequireditems

func minimumCost(cost1 int, cost2 int, costBoth int, need1 int, need2 int) int64 {
	cost1i64 := int64(cost1)
	cost2i64 := int64(cost2)
	costBothi64 := int64(costBoth)
	need1i64 := int64(need1)
	need2i64 := int64(need2)


	individual := cost1i64*need1i64 + cost2i64*need2i64
	combined := costBothi64 * max(need1i64, need2i64)
	m := min(need1i64, need2i64)
	partial := (costBothi64 * m) + ((need1i64 - m) * cost1i64) + ((need2i64 - m) * cost2i64)

	return min(individual, combined, partial)
}
