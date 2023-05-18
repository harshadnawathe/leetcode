package nqueens

import (
	"bytes"
	"unsafe"
)

func solveNQueens(n int) [][]string {
	boards := []board{}
	var f func(board)
	f = func(b board) {
		if isSolved(b) {
			boards = append(boards, b)
			return
		}

		for _, next := range placeNextQueen(b) {
			f(next)
		}
	}
	f(newBoard(n))

	result := [][]string{}
	for _, b := range boards {
		result = append(result, asStrSlice(b))
	}
	return result
}

type board struct {
	size       int
	placements []int
}

func newBoard(n int) board {
	return board{n, nil}
}

func isSolved(b board) bool {
	return b.size == len(b.placements)
}

func clone(b board) board {
	c := newBoard(b.size)
	c.placements = make([]int, len(b.placements))
	copy(c.placements, b.placements)
	return c
}

func availablePositions(b board) map[int]struct{} {
	positions := make(map[int]struct{})
	for i := 0; i < b.size; i++ {
		positions[i] = struct{}{}
	}
	row := len(b.placements)
	for i, q := range b.placements {
		delete(positions, q)
		depth := row - i
		delete(positions, q-depth)
		delete(positions, q+depth)
	}

	return positions
}

func placeNextQueen(b board) []board {
	options := []board{}
	for pos := range availablePositions(b) {
		c := clone(b)
		c.placements = append(c.placements, pos)
		options = append(options, c)
	}
	return options
}

func asStrSlice(b board) []string {
	slice := []string{}
	for _, q := range b.placements {
		row := bytes.Repeat([]byte{'.'}, b.size)
		row[q] = 'Q'
		slice = append(slice, *(*string)(unsafe.Pointer(&row)))
	}
	return slice
}
