package clonegraph

import (
	"fmt"
	"sort"

	"github.com/google/go-cmp/cmp"
	"github.com/google/go-cmp/cmp/cmpopts"
)

type Node struct {
	Val       int
	Neighbors []*Node
}

func (g *Node) String() string {
	return fmt.Sprintf("%+v", adjListOf(g))
}

func graphOf(adjList [][]int) *Node {
	if len(adjList) == 0 {
		return nil
	}

	nodes := make([]*Node, len(adjList))

	for i, neighbours := range adjList {
		if nodes[i] == nil {
			nodes[i] = &Node{Val: i + 1}
		}

		for _, neighbour := range neighbours {
			if nodes[neighbour-1] == nil {
				nodes[neighbour-1] = &Node{Val: neighbour}
			}

			nodes[i].Neighbors = append(nodes[i].Neighbors, nodes[neighbour-1])
		}
	}

	return nodes[0]
}

func adjListOf(graph *Node) [][]int {
	if graph == nil {
		return nil
	}

	node := make(map[int][]int)
	visited := make(map[int]struct{})

	q := []*Node{graph}
	add := func(n *Node) { q = append(q, n) }
	remove := func() { q = q[1:] }
	next := func() *Node { return q[0] }

	for len(q) > 0 {
		nxt := next()
		remove()

		if _, found := visited[nxt.Val]; found {
			continue
		}

		for _, neighbour := range nxt.Neighbors {
			node[nxt.Val] = append(node[nxt.Val], neighbour.Val)
			add(neighbour)
		}
		visited[nxt.Val] = struct{}{}
	}

	adjList := [][]int{}
	for _, neighbours := range node {
		adjList = append(adjList, neighbours)
	}
	return adjList
}

func equal(g1, g2 *Node) bool {
	l1 := adjListOf(g1)
	for _, a := range l1 {
		sort.Ints(a)
	}

	l2 := adjListOf(g2)
	for _, a := range l2 {
		sort.Ints(a)
	}

	return cmp.Equal(l1, l2, cmpopts.SortSlices(func(a, b []int) bool {
		if len(a) == 0 || len(b) == 0 {
			return len(a) < len(b)
		}

		return a[0] < b[0]
	}))
}
