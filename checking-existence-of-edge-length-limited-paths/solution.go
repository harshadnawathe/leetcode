package checkingexistenceofedgelengthlimitedpaths

import "sort"

func distanceLimitedPathsExist(n int, edgeList [][]int, queries [][]int) []bool {
	for i := range queries {
		queries[i] = append(queries[i], i)
	}

	sort.Slice(queries, func(i, j int) bool {
		return queries[i][2] < queries[j][2]
	})

	sort.Slice(edgeList, func(i, j int) bool {
		return edgeList[i][2] < edgeList[j][2]
	})

	ds := newDisjoinSets(n)
	pathExists := make([]bool, len(queries))
	edgeIdx := 0
	for _, query := range queries {
		p, q, limit, idx := query[0], query[1], query[2], query[3]

		for edgeIdx < len(edgeList) && edgeList[edgeIdx][2] < limit {
			n1 := edgeList[edgeIdx][0]
			n2 := edgeList[edgeIdx][1]
			join(ds, n1, n2)
			edgeIdx++
		}

		pathExists[idx] = connected(ds, p, q)
	}

	return pathExists
}

type disjointSets []int

func newDisjoinSets(n int) disjointSets {
	ds := make([]int, 0, n)
	for i := 0; i < n; i++ {
		ds = append(ds, -1)
	}
	return ds
}

func parent(ds disjointSets, node int) int {
	if ds[node] >= 0 {
		return parent(ds, ds[node])
	}
	return node
}

func rank(ds disjointSets, node int) int {
	p := parent(ds, node)
	return -ds[p]
}

func connected(ds disjointSets, node1, node2 int) bool {
	return parent(ds, node1) == parent(ds, node2)
}

func join(ds disjointSets, node1, node2 int) {
	parent1, parent2 := parent(ds, node1), parent(ds, node2)

	if parent1 == parent2 {
		return
	}

	rank1, rank2 := rank(ds, parent1), rank(ds, parent2)
	if rank1 < rank2 {
		ds[parent2] -= rank1
		ds[parent1] = parent2
	} else {
		ds[parent1] -= rank2
		ds[parent2] = parent1
	}
}
