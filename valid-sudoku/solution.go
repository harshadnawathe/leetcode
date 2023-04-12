package validsudoku

const empty = '.'

const (
	regionNotSet = iota
	regionRow
	regionColumn
	regionGrid
)

type valueEntry struct {
	regionType, region uint8
	value              byte
}

func isValidSudoku(board [][]byte) bool {
	values := make(map[valueEntry]struct{})

	for r, row := range board {
		for c, value := range row {
			if value == empty {
				continue
			}

			// row entry
			rowEntry := valueEntry{regionRow, uint8(r), value}
			if _, found := values[rowEntry]; found {
				return false
			}
			values[rowEntry] = struct{}{}

			// col entry
			colEntry := valueEntry{regionColumn, uint8(c), value}
			if _, found := values[colEntry]; found {
				return false
			}
			values[colEntry] = struct{}{}

			// grid entry
			gridEntry := valueEntry{
				regionType: regionGrid,
				region:     uint8(3*(r/3) + (c / 3)),
				value:      value,
			}
			if _, found := values[gridEntry]; found {
				return false
			}
			values[gridEntry] = struct{}{}
		}
	}
	return true
}
