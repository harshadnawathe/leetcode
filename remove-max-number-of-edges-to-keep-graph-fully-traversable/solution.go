package removemaxnumberofedgestokeepgraphfullytraversable

const (
	edgeAlice = 1
	edgeBob   = 2
	edgeAll   = 3
)

func maxNumEdgesToRemove(n int, edges [][]int) int {
	alice := newDisjointsSet(n)
	bob := newDisjointsSet(n)

	numEdgesToBeRemoved := 0
	for _, edge := range edges {
		if edge[0] != edgeAll {
			continue
		}
		if !alice.Join(edge[1], edge[2]) || !bob.Join(edge[1], edge[2]) {
			numEdgesToBeRemoved++
		}
	}

	for _, edge := range edges {
		if edge[0] == edgeAlice {
			if !alice.Join(edge[1], edge[2]) {
				numEdgesToBeRemoved++
			}
		} else if edge[0] == edgeBob {
			if !bob.Join(edge[1], edge[2]) {
				numEdgesToBeRemoved++
			}
		}
	}
	if !alice.IsConnected() || !bob.IsConnected() {
		return -1
	}

	return numEdgesToBeRemoved
}

type disjointsSet struct {
	s []int
	n int
}

func newDisjointsSet(n int) *disjointsSet {
	ds := &disjointsSet{
		s: make([]int, 0, n+1),
		n: n,
	}
	for i := 0; i < n+1; i++ {
		ds.s = append(ds.s, -1)
	}
	return ds
}

func parent(ds *disjointsSet, node int) int {
	if ds.s[node] >= 0 {
		return parent(ds, ds.s[node])
	}
	return node
}

func rank(ds *disjointsSet, node int) int {
	p := parent(ds, node)
	return -ds.s[p]
}

func (ds *disjointsSet) Len() int { return ds.n }
func (ds *disjointsSet) Join(node1, node2 int) bool {
	parent1, parent2 := parent(ds, node1), parent(ds, node2)
	if parent1 == parent2 {
		return false
	}

	ds.n--
	rank1, rank2 := rank(ds, node1), rank(ds, node2)
	if rank1 < rank2 {
		ds.s[parent1] = parent2
		ds.s[parent2] -= rank1
	} else {
		ds.s[parent2] = parent1
		ds.s[parent1] -= rank2
	}
	return true
}

func (ds *disjointsSet) IsConnected() bool {
	return ds.n == 1
}
