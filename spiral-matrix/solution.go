package spiralmatrix

import (
	"errors"
	"math"
)

// visited value indicates element is already collected
// Matrix contains values between -100 to 100
const visited int = math.MinInt32

type direction int

const (
	directionRight direction = iota
	directionDown
	directionLeft
	directionUp
)

func (d direction) Next(row, col int) (int, int) {
	switch d {
	case directionRight:
		col++
	case directionDown:
		row++
	case directionLeft:
		col--
	case directionUp:
		row--
	}
	return row, col
}

func (d direction) Turn() direction {
	switch d {
	case directionRight:
		return directionDown
	case directionDown:
		return directionLeft
	case directionLeft:
		return directionUp
	case directionUp:
		return directionRight
	}
	panic("invalid direction")
}

type spiralIterator struct {
	mat      [][]int
	row, col int
	dir      direction
	err      error
}

func newSpiralIterator(mat [][]int) *spiralIterator {
	itr := &spiralIterator{mat, 0, 0, directionRight, nil}
	if !canVisit(itr.mat, 0, 0) {
		itr.err = errors.New("EOF")
	}
	return itr
}

func canVisit(mat [][]int, row, col int) bool {
	return row >= 0 && row < len(mat) && col >= 0 && col < len(mat[0]) && mat[row][col] != visited
}

func (itr *spiralIterator) Next() {
	row, col := itr.dir.Next(itr.row, itr.col)
	if !canVisit(itr.mat, row, col) {
		itr.dir = itr.dir.Turn()
		row, col = itr.dir.Next(itr.row, itr.col)
		if !canVisit(itr.mat, row, col) {
			itr.err = errors.New("EOF")
			return
		}
	}
	itr.mat[itr.row][itr.col] = visited
	itr.row, itr.col = row, col
}

func (itr *spiralIterator) HasValue() bool {
	return itr.err == nil
}

func (itr *spiralIterator) Value() int {
	return itr.mat[itr.row][itr.col]
}

func spiralOrder(matrix [][]int) []int {
	var result []int
	for itr := newSpiralIterator(matrix); itr.HasValue(); itr.Next() {
		result = append(result, itr.Value())
	}
	return result
}
