package minimumgeneticmutation

func minMutation(startGene string, endGene string, bank []string) int {
	validGenes := make(map[string]struct{})
	for _, gene := range bank {
		validGenes[gene] = struct{}{}
	}

	nextBase := map[byte][]byte{
		'A': {'C', 'G', 'T'},
		'C': {'A', 'G', 'T'},
		'G': {'A', 'C', 'T'},
		'T': {'A', 'C', 'G'},
	}

	type geneGeneration struct {
		gene       string
		generation int
	}

	visited := make(map[string]struct{})

	q := []geneGeneration{{startGene, 0}}

	for len(q) > 0 {
		g := q[0]
		q = q[1:]

		if _, ok := visited[g.gene]; ok {
			continue
		}

		visited[g.gene] = struct{}{}

		if g.gene == endGene {
			return g.generation
		}

		g.generation++

		for i := 0; i < len(g.gene); i++ {
			for _, base := range nextBase[g.gene[i]] {
				next := g.gene[:i] + string(base) + g.gene[i+1:]
				if _, ok := validGenes[next]; ok {
					q = append(q, geneGeneration{next, g.generation})
				}
			}
		}

	}

	return -1
}
