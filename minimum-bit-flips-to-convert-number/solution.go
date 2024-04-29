package minimumbitflipstoconvertnumber

func minBitFlips(start int, goal int) int {
	flipCount := 0
	for start > 0 || goal > 0 {
		if start&0x1 != goal&0x1 {
			flipCount++
		}
		start >>= 1
		goal >>= 1
	}
	return flipCount
}
