package designafoodratingsystem

import (
	"container/heap"
	"strings"
)

type FoodRatings struct {
	foodCusine  map[string]string
	cuisineHeap map[string]*cuisineHeap
}

func Constructor(foods []string, cuisines []string, ratings []int) FoodRatings {
	fr := FoodRatings{
		make(map[string]string),
		make(map[string]*cuisineHeap),
	}

	for i := 0; i < len(foods); i++ {
		food, cuisine, rating := foods[i], cuisines[i], ratings[i]
		fr.foodCusine[food] = cuisine
		h, ok := fr.cuisineHeap[cuisine]
		if !ok {
			h = &cuisineHeap{
				foodHeapPos: make(map[string]int),
			}
			heap.Init(h)
			fr.cuisineHeap[cuisine] = h
		}
		h.ChangeRating(food, rating)
	}

	return fr
}

func (fr *FoodRatings) ChangeRating(food string, newRating int) {
	cuisine := fr.foodCusine[food]
	fr.cuisineHeap[cuisine].ChangeRating(food, newRating)
}

func (fr *FoodRatings) HighestRated(cuisine string) string {
	return fr.cuisineHeap[cuisine].heap[0].Name
}

type foodRating struct {
	Name   string
	Rating int
}

type cuisineHeap struct {
	foodHeapPos map[string]int
	heap        []*foodRating
}

func (h *cuisineHeap) ChangeRating(food string, newRating int) {
	if pos, ok := h.foodHeapPos[food]; ok {
		h.heap[pos].Rating = newRating
		heap.Fix(h, pos)
	} else {
		heap.Push(h, &foodRating{food, newRating})
	}
}

// Len is the number of elements in the collection.
func (h *cuisineHeap) Len() int {
	return len(h.heap)
}

// Less reports whether the element with index i
// must sort before the element with index j.
//
// If both Less(i, j) and Less(j, i) are false,
// then the elements at index i and j are considered equal.
// Sort may place equal elements in any order in the final result,
// while Stable preserves the original input order of equal elements.
//
// Less must describe a transitive ordering:
//   - if both Less(i, j) and Less(j, k) are true, then Less(i, k) must be true as well.
//   - if both Less(i, j) and Less(j, k) are false, then Less(i, k) must be false as well.
//
// Note that floating-point comparison (the < operator on float32 or float64 values)
// is not a transitive ordering when not-a-number (NaN) values are involved.
// See Float64Slice.Less for a correct implementation for floating-point values.
func (h *cuisineHeap) Less(i int, j int) bool {
	if h.heap[i].Rating == h.heap[j].Rating {
		return strings.Compare(h.heap[i].Name, h.heap[j].Name) < 0
	}
	return h.heap[i].Rating > h.heap[j].Rating
}

// Swap swaps the elements with indexes i and j.
func (h *cuisineHeap) Swap(i int, j int) {
	h.foodHeapPos[h.heap[i].Name] = j
	h.foodHeapPos[h.heap[j].Name] = i
	h.heap[i], h.heap[j] = h.heap[j], h.heap[i]
}

func (h *cuisineHeap) Push(x any) {
	fr := x.(*foodRating)
	h.heap = append(h.heap, fr)
	h.foodHeapPos[fr.Name] = h.Len() - 1
}

func (h *cuisineHeap) Pop() any {
	old := h.heap
	n := len(old)
	x := old[n-1]
	h.heap = old[0 : n-1]
	delete(h.foodHeapPos, x.Name)
	return x
}
