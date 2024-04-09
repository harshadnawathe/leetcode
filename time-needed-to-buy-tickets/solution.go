package timeneededtobuytickets

func timeRequiredToBuy(tickets []int, k int) int {
	kTicketCount := tickets[k]

	time := 0

	for i, ticket := range tickets {
		if i < k {
			if ticket >= kTicketCount {
				time += kTicketCount
			} else {
				time += ticket
			}
		} else if i > k {
			if ticket >= kTicketCount {
				time += kTicketCount - 1
			} else {
				time += ticket
			}
		} else {
			time += kTicketCount
		}
	}

	return time
}
