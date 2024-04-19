package numberoflaserbeamsinabank

func numberOfBeams(bank []string) int {
	numBeams := 0

	numBeamsInPrevRow := 0

	for _, row := range bank {
		numBeamsInCurrRow := 0
		for _, cell := range row {
			if cell == '1' {
				numBeamsInCurrRow++
			}
		}

		if numBeamsInCurrRow != 0 {
			numBeams += numBeamsInPrevRow * numBeamsInCurrRow
			numBeamsInPrevRow = numBeamsInCurrRow
		}
	}

	return numBeams
}
